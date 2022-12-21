package controller

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
	"vitalic_project/internal/app/repository"
)

type Controller struct {
	baseVariableController *BaseVariableController
	termodPropController   *TermodPropController
	potentialController    *PotentialController
	balanceController      *BalanceController
	instabilityController  *InstabilityController
}

func NewControllers() *Controller {
	return &Controller{
		baseVariableController: NewBaseVariableController(),
		termodPropController:   NewTermodPropController(),
		potentialController:    NewPotentialController(),
		balanceController:      NewBalanceController(),
		instabilityController:  NewInstabilityController(),
	}
}

func (controller *Controller) BuildControllers(repo *repository.MainRepo, router *mux.Router, fill bool, logger *logrus.Logger) {
	controller.OpenMainPage(router)
	controller.OpenJsFilePage(router)
	controller.OpenSendPage(router)
	controller.OpenSendJs(router)
	if fill {
		logger.Info("Начинаем заполнять базу данных автоматический...")
		FillingDatabase(repo)
	}
	controller.instabilityController.CreateInstabilityController(repo.InstabilityRepo, router)
	controller.termodPropController.CreateThePrElController(repo.TermodPropRepo, router)
	controller.baseVariableController.CreateBaseVrElController(repo.BaseVariableRepo, router)
	controller.potentialController.CreatePotentialController(repo.PotentialsRepo, router)
	controller.balanceController.CreateBalanceController(repo.BalanceRepo, router)

	controller.instabilityController.FindInstabilityElController(repo.InstabilityRepo, router)
	controller.termodPropController.FindThermoChElController(repo.TermodPropRepo, router)
	controller.baseVariableController.FindBaseVariableController(repo.BaseVariableRepo, router)
	controller.potentialController.FindPotentialElController(repo.PotentialsRepo, router)
	controller.balanceController.FindBalanceElController(repo.BalanceRepo, router)
}

func (controller *Controller) OpenMainPage(router *mux.Router) *mux.Route {
	return router.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "internal/app/frontend/main.html")
	})
}
func (controller *Controller) OpenJsFilePage(router *mux.Router) *mux.Route {
	return router.HandleFunc("/main.js", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "internal/app/frontend/main.js")
	})
}

func (controller *Controller) OpenSendPage(router *mux.Router) *mux.Route {
	return router.HandleFunc("/send", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "internal/app/frontend/send.html")
	})
}
func (controller *Controller) OpenSendJs(router *mux.Router) *mux.Route {
	return router.HandleFunc("/send.js", func(writer http.ResponseWriter, request *http.Request) {
		http.ServeFile(writer, request, "internal/app/frontend/send.js")
	})
}
