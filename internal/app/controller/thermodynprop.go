package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/repository"
)

type TermodPropController struct {
	Router *mux.Router
}

func NewTermodPropController() *TermodPropController {
	return &TermodPropController{}
}
func (tpc *TermodPropController) CreateThePrElController(r *repository.TermodPropRepo, rou *mux.Router) {
	tpc.Router = rou
	tpc.Router.HandleFunc("/chemistry/create/term", func(writer http.ResponseWriter, request *http.Request) {
		m := model.TermodProp{}
		json.NewDecoder(request.Body).Decode(&m)
		fmt.Println(m)
		readParam := ReadAllThermoParam(&m)
		if readParam == "" {
			if _, err := r.Create(&m); err != nil {
				fmt.Println(err)
				fmt.Fprintf(writer, "Не удалось создать элемент. Возможно он уже есть")
			} else {
				fmt.Fprintf(writer, "Элемент успешно добавлен!")
			}
		} else if readParam != "" {
			fmt.Fprintf(writer, readParam)
		}
	})
}

func (tpc *TermodPropController) FindThermoChElByEmailController(r *repository.TermodPropRepo, rou *mux.Router) {
	tpc.Router = rou
	tpc.Router.HandleFunc("/chemistry/term", func(writer http.ResponseWriter, request *http.Request) {
		m, err := r.FindTherByName(request.URL.Query().Get("name"))
		if err == sql.ErrNoRows {
			fmt.Fprintf(writer, "Такого элемента нет!")
		} else if m != nil && err == nil {
			returnedModel, err := json.Marshal(m)
			if err != nil {
				fmt.Fprintf(writer, "Неудалось распокавать json")
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(returnedModel)
		}
	})
}
func ReadAllThermoParam(mod *model.TermodProp) string {
	if mod.ElName == "" {
		return "нет имени элемента или оно пустое"
	}
	if mod.FirstParam == "" {
		return "нет первого параметра термодинамических свойств элемента"
	}
	if mod.SecondParam == "" {
		return "нет второго параметра термодинамических свойств элемента"
	}
	if mod.ThirdParam == "" {
		return "нет третьего параметра термодинамических свойств элемента"
	}
	if mod.FourthParam == "" {
		return "нет четвертого параметра термодинамических свойств элемента"
	}
	return ""
}
