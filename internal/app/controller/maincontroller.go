package controller

import (
	"github.com/gorilla/mux"
	"vitalic_project/internal/app/repository"
)

type Controller struct {
	baseVariableController *BaseVariableController
	termodPropController   *TermodPropController
}

func NewControllers() *Controller {
	return &Controller{
		baseVariableController: NewBaseVariableController(),
		termodPropController:   NewTermodPropController(),
	}
}

func (controller *Controller) BuildControllers(repo *repository.MainRepo, router *mux.Router) {
	controller.termodPropController.CreateThePrElController(repo.TermodPropRepo, router)
	controller.termodPropController.FindThermoChElByEmailController(repo.TermodPropRepo, router)
	controller.baseVariableController.FindBaseVariableElByEmailController(repo.BaseVariableRepo, router)
	controller.baseVariableController.CreateBaseVrElController(repo.BaseVariableRepo, router)

}
