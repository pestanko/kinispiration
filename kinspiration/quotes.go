package kinspiration

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	ID     string `json:"id,omitempty"`
	Body   string `json:"body,omitempty"`
	Author string `json:"author,omitempty"`
}

type Quotes struct {
	Collection []Quote
	App        *App
}

func (quotes *Quotes) FilePath() string {
	return quotes.App.Config.FilePath
}

func (quotes *Quotes) Init(app *App) {
	quotes.App = app
	quotes.ReadAllQuotes()
	unix := time.Now().Unix()
	rand.Seed(unix)
}

func (quotes *Quotes) ReadAllQuotes() {
	jsonFile, err := os.Open(quotes.FilePath())
	if err != nil {
		log.Printf("%s", err)
		return
	}
	defer jsonFile.Close()
	byteValues, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteValues, &quotes.Collection)
}

func (quotes *Quotes) WriteAllQuotes() {
	marshal, _ := json.Marshal(quotes.Collection)
	ioutil.WriteFile(quotes.FilePath(), marshal, 0644)
}

func (quotes *Quotes) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	json.NewEncoder(w).Encode(quotes.Collection)
}

func (quotes *Quotes) CreateQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)
	quote.ID = params["id"]
	go quotes.AddQuote(quote)
	w.WriteHeader(204)
}

func (quotes *Quotes) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	randomQuote := quotes.Collection[rand.Intn(len(quotes.Collection))]
	json.NewEncoder(w).Encode(randomQuote)
}

func (quotes *Quotes) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	go quotes.DelQuote(params["id"])
	w.WriteHeader(204)
}

func (quotes *Quotes) RegisterQuotes() {
	quotes.App.Get("/quotes", quotes.GetAllQuotes)
	quotes.App.Post("/quotes/{id}", quotes.CreateQuote)
	quotes.App.Delete("/quotes/{id}", quotes.DeleteQuote)
	quotes.App.Get("/random", quotes.GetRandomQuote)
}

func (quotes *Quotes) AddQuote(quote Quote) {
	quotes.Collection = append(quotes.Collection, quote)
	quotes.WriteAllQuotes()
}

func (quotes *Quotes) DelQuote(id string) {
	for i, q := range quotes.Collection {
		if q.ID == id {
			quotes.Collection = append(quotes.Collection[:i], quotes.Collection[i+1:]...)
		}
	}
	quotes.WriteAllQuotes()
}
