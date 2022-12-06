package apiserver

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	store2 "vitalic_project/internal/app/store"
)

type APIserver struct {
	config *ServerConfig
	logger *logrus.Logger
	router *mux.Router
	store  *store2.Store
}

func ServerInit(config *ServerConfig) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store2.NewStore(store2.NewDataBaseConfig()),
	}
}

func (api *APIserver) Start() error {
	if err := api.loggerConfig(); err != nil {
		return err
	}
	api.logger.Info("Сосать жопу")

	api.routerConfig()

	if err := api.store.Open(); err != nil {
		return err
	}

	return http.ListenAndServe("localhost"+api.config.Bind, api.router)
}

func (api *APIserver) loggerConfig() error {
	level, err := logrus.ParseLevel(api.config.LoggerLevel)
	if err != nil {
		return err
	}
	api.logger.SetLevel(level)
	return nil
}

func (api *APIserver) routerConfig() {
	conn := api.router
	conn.HandleFunc("/chemistry", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintf(writer, "Жопа хуй")
	})
}
