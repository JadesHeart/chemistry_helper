package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"vitalic_project/internal/app/controller"
	"vitalic_project/internal/app/repository"
	"vitalic_project/internal/app/store"
)

type APIserver struct {
	config *ServerConfig
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
	repo   *repository.MainRepo
}

func ServerInit(config *ServerConfig) *APIserver {
	return &APIserver{
		config: config,
		logger: logrus.New(),
		router: mux.NewRouter(),
		store:  store.NewStore(store.NewDataBaseConfig()),
	}
}

func (api *APIserver) Start() error {
	if err := api.loggerConfig(); err != nil {
		return err
	}
	api.logger.Info("Сервер запустился...")

	if err := api.store.Open(); err != nil {
		return err
	}

	api.repoConfig()

	api.routerConfig()

	return http.ListenAndServe("localhost"+api.config.Bind, api.router)
}

func (api *APIserver) repoConfig() {
	api.repo = repository.NewMainRepo(api.store)
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
	contr := controller.NewControllers()
	contr.BuildControllers(api.repo, api.router)
}
