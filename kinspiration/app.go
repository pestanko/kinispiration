package kinspiration

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	Quotes Quotes
	Config *Config
	AuthMiddle AuthMiddleware
}

// Init initializes the app with predefined configuration
func (a *App) Init(config *Config) {
	a.Router = mux.NewRouter()
	a.Config = config
	a.Quotes = Quotes{}
	a.Quotes.Init(a)
	a.setRouters()
}

// setRouters sets the all required routers
func (a *App) setRouters() {
	// Routing for handling the projects
	a.Quotes.RegisterQuotes()
	a.AuthMiddle = AuthMiddleware{App: a}
	a.Router.Use(a.AuthMiddle.Middleware)
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	logHandler := a.handleLogFunc(f)
	a.Router.Handle(path, logHandler).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	logHandler := a.handleLogFunc(f)
	a.Router.Handle(path, logHandler).Methods("POST")
}

// Put wraps the router for PUT method
func (a *App) Put(path string, f func(w http.ResponseWriter, r *http.Request)) {
	logHandler := a.handleLogFunc(f)
	a.Router.Handle(path, logHandler).Methods("PUT")
}

// Delete wraps the router for DELETE method
func (a *App) Delete(path string, f func(w http.ResponseWriter, r *http.Request)) *mux.Route {
	logHandler := a.handleLogFunc(f)
	return a.Router.Handle(path, logHandler).Methods("DELETE")
}

func (a *App) handleLogFunc(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	funcHandler := http.HandlerFunc(f)
	return  handlers.CombinedLoggingHandler(os.Stderr, funcHandler)
}

func (a *App) Run() {
	log.Fatal(http.ListenAndServe(a.Config.Host, a.Router))
}
