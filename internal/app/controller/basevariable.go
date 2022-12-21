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

type BaseVariableController struct {
	Router *mux.Router
}

func NewBaseVariableController() *BaseVariableController {
	return &BaseVariableController{}
}
func (inc *BaseVariableController) CreateBaseVrElController(r *repository.BaseVariableRepo, rou *mux.Router) {
	inc.Router = rou
	inc.Router.HandleFunc("/chemistry/create/base", func(writer http.ResponseWriter, request *http.Request) {
		m := model.BaseVariable{}
		json.NewDecoder(request.Body).Decode(&m)
		readParam := ReadAllParam(&m)
		if readParam == "" {
			if _, err := r.Create(&m); err != nil {
				fmt.Fprintf(writer, "Не удалось создать элемент. Возможно он уже есть или данные не валидны!")
			} else {
				fmt.Fprintf(writer, "Элемент успешно добавлен!")
			}
		} else if readParam != "" {
			fmt.Fprintf(writer, readParam)
		}
	})
}

func (inc *BaseVariableController) FindBaseVariableController(r *repository.BaseVariableRepo, rou *mux.Router) {
	inc.Router = rou
	inc.Router.HandleFunc("/chemistry/base", func(writer http.ResponseWriter, request *http.Request) {
		m, err := r.FindBaseVarElByName(request.URL.Query().Get("name"))
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
func ReadAllParam(mod *model.BaseVariable) string {
	if mod.ElName == "" {
		return "нет имени элемента или оно пустое"
	}
	if mod.ElectronicConfiguration == "" {
		return "нет электронной конфигурации или она пустая"
	}
	if mod.StabilityOxidation == "" {
		return "нет устойчивость степени окисления и их конфигурации или оно пустое"
	}
	if mod.MeltingPoint == "" {
		return "нет температуры плавления или она пустая"
	}
	if mod.BoilingPoint == "" {
		return "нет температуры кипения или она пустая"
	}
	if mod.ChemicalCompounds == "" {
		return "нет химических соединений или они пустые"
	}
	return ""
}
