package kinspiration

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

type Quote struct {
	ID     string `json:"id,omitempty"`
	Body   string `json:"quote,omitempty"`
	Author string `json:"author,omitempty"`
}

type Quotes struct {
	Collection []Quote
	App        *App
}

func JSONWrite(w io.Writer, obj interface{}) {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	encoder.Encode(obj)
}

func (quotes *Quotes) FilePath() string {
	return quotes.App.Config.QuotesPath
}

func (quotes *Quotes) Init(app *App) {
	quotes.App = app
	quotes.ReadAllQuotes()
	unix := time.Now().Unix()
	rand.Seed(unix)
}

func (quotes *Quotes) ReadAllQuotes() {
	log.Printf("[QUOTES] Loading all quotes from: %s", quotes.FilePath())
	os.OpenFile(quotes.FilePath(), os.O_RDONLY|os.O_CREATE, 0666)
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
	log.Printf("[QUOTES] Writing all quotes to: %s", quotes.FilePath())
	marshal, _ := json.MarshalIndent(quotes.Collection, "", "  ")
	ioutil.WriteFile(quotes.FilePath(), marshal, 0644)
}


func (quotes *Quotes) GetAllQuotes(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	JSONWrite(w, quotes.Collection)
}


func (quotes *Quotes) CreateQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)
	quotes.AddQuote(quote, params["id"], true)
	JSONWrite(w, quote)
}

func (quotes *Quotes) CreateQuoteRandom(w http.ResponseWriter, r *http.Request) {
	var quote Quote
	_ = json.NewDecoder(r.Body).Decode(&quote)
	quotes.AddQuote(quote, "", true)
	JSONWrite(w, quote)
}

func (quotes *Quotes) GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	quote := quotes.Collection[rand.Intn(len(quotes.Collection))]
	JSONWrite(w, quote)
}

func (quotes *Quotes) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	quotes.DelQuote(params["id"])
	w.WriteHeader(204)
}

func (quotes *Quotes) ImportQuotes(w http.ResponseWriter, r *http.Request) {
	var quoteList []Quote
	_ = json.NewDecoder(r.Body).Decode(&quoteList)

	for _, quote := range quoteList {
		quotes.AddQuote(quote, "", false)
	}

	quotes.WriteAllQuotes()
	w.WriteHeader(200)
	JSONWrite(w, quotes.Collection)
}

func (quotes *Quotes) RegisterQuotes() {
	quotes.App.Get("/quotes", quotes.GetAllQuotes)
	quotes.App.Post("/quotes/{id}", quotes.CreateQuote)
	quotes.App.Post("/quotes", quotes.CreateQuoteRandom)
	quotes.App.Delete("/quotes/{id}", quotes.DeleteQuote)
	quotes.App.Get("/random", quotes.GetRandomQuote)
	quotes.App.Post("/import", quotes.ImportQuotes)
}

func (quotes *Quotes) AddQuote(quote Quote, id string, save bool) {
	if id == "" {
		id = uuid.New().String()
	}
	quote.ID = id
	quotes.Collection = append(quotes.Collection, quote)
	if save {
		quotes.WriteAllQuotes()
	}
}

func (quotes *Quotes) DelQuote(id string) {
	for i, q := range quotes.Collection {
		if q.ID == id {
			quotes.Collection = append(quotes.Collection[:i], quotes.Collection[i+1:]...)
		}
	}
	quotes.WriteAllQuotes()
}
