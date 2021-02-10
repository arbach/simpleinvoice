package app

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/arbach/simpleinvoice/app/handlers"
	"github.com/arbach/simpleinvoice/app/services"
	"github.com/arbach/simpleinvoice/app/store/memstore"
	"github.com/arbach/simpleinvoice/app/store/sqlstore"
	"github.com/arbach/simpleinvoice/ethclient"

	"github.com/jmoiron/sqlx"
)

var AppInstance *App

type App struct {
	Router    *mux.Router
	DB        *sqlx.DB
	handlers  *handlers.Handler
	EthClient *ethclient.Client
}

func (app *App) SetupRoutes() {
	AppInstance = app
	app.Router = mux.NewRouter()
	app.pingRoute()

	service := services.New(app.EthClient)
	if os.Getenv("USE_DB_STORAGE") != "true" {
		store := memstore.New()
		app.handlers = handlers.New(store, service)
	} else {
		store := sqlstore.New(app.DB)
		app.handlers = handlers.New(store, service)
	}

	app.setupRouteHandlers()
}

func (app *App) setupRouteHandlers() {
	app.Router.HandleFunc("/invoice", app.handlers.GetInvoice).Methods("GET")
	app.Router.HandleFunc("/invoice", app.handlers.GenerateInvoice).Methods("POST")
}

func (app *App) pingRoute() {
	app.Router.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "pong")
	})
}
