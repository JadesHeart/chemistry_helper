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

type InstabilityController struct {
	Router *mux.Router
}

func NewInstabilityController() *InstabilityController {
	return &InstabilityController{}
}
func (inc *InstabilityController) CreateInstabilityController(r *repository.InstabilityRepo, rou *mux.Router) {
	inc.Router = rou
	inc.Router.HandleFunc("/chemistry/create/instability", func(writer http.ResponseWriter, request *http.Request) {
		m := model.Instability{}
		json.NewDecoder(request.Body).Decode(&m)
		readParam := ReadAllInstabilityParam(&m)
		if readParam == "" {
			if _, err := r.CreateEl(&m); err != nil {
				fmt.Fprintf(writer, "Не удалось создать элемент. Возможно он уже есть или данные не валидны!")
			} else {
				fmt.Fprintf(writer, "Элемент успешно добавлен!")
			}
		} else if readParam != "" {
			fmt.Fprintf(writer, readParam)
		}
	})
}

func (inc *InstabilityController) FindInstabilityElController(r *repository.InstabilityRepo, rou *mux.Router) {
	inc.Router = rou
	inc.Router.HandleFunc("/chemistry/instability", func(writer http.ResponseWriter, request *http.Request) {
		m, err := r.FindInstabilityElByName(request.URL.Query().Get("name"))
		if err == sql.ErrNoRows || len(m) == 0 {
			fmt.Fprintf(writer, "Такого или таких элемента(ов) нет!")
		} else if m != nil && err == nil && len(m) != 0 {
			returnedModel, err := json.Marshal(m)
			if err != nil {
				fmt.Fprintf(writer, "Вот тут ошиб04ка вышла")
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(returnedModel)
		}
	})
}
func ReadAllInstabilityParam(mod *model.Instability) string {
	if mod.ElName == "" {
		return "нет имени элемента или оно пустое"
	}
	if mod.Ligand == "" {
		return "нет лиганда"
	}
	if mod.Complex == "" {
		return "нет комплекса"
	}
	if mod.LastParam == "" {
		return "нет К(нест)"
	}
	return ""
}
