package main

import (
	"fmt"

	"net/http"
	"time"

	"github.com/arbach/simpleinvoice/app"
	"github.com/arbach/simpleinvoice/db"
	"github.com/arbach/simpleinvoice/ethclient"
	"github.com/kelseyhightower/envconfig"
	log "github.com/sirupsen/logrus"
	"github.com/subosito/gotenv"
)

type ServiceArgs struct {
	HostPort       int    `envconfig:"HOST_PORT" default:"8080"`
	NodeURL        string `envconfig:"ETHEREUM_NODE_URL" default:"http://localhost:8445"`
	UseDbStorage   bool   `envconfig:"USE_DB_STORAGE" default:"false"`
	DatabaseConfig db.DatabaseConfig
}

func main() {
	gotenv.Load("config.env")
	cfg := ServiceArgs{}
	if err := envconfig.Process("", &cfg); err != nil {
		panic(err)
	}
	startApi(cfg)
}

func startApi(cfg ServiceArgs) {
	app := app.App{}

	if cfg.UseDbStorage {
		log.Info("Initializing DB")
		app.DB = db.SetupSqlxDB(cfg.DatabaseConfig)
	}

	var err error
	log.Info("Initializing Ethereum Client")
	app.EthClient, err = ethclient.NewClient(cfg.NodeURL)
	if err != nil {
		panic(err)
	}

	log.Info("Setting api routes")
	app.SetupRoutes()
	http.Handle("/", app.Router)
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", cfg.HostPort),
		Handler:        nil,
		ReadTimeout:    2 * time.Second,
		WriteTimeout:   2 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Info(fmt.Sprintf("App running at: %s", s.Addr))
	log.Fatal(s.ListenAndServe())
}
