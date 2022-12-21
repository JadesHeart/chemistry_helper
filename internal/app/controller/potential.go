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

type PotentialController struct {
	Router *mux.Router
}

func NewPotentialController() *PotentialController {
	return &PotentialController{}
}
func (pt *PotentialController) CreatePotentialController(r *repository.PotentialsRepo, rou *mux.Router) {
	pt.Router = rou
	pt.Router.HandleFunc("/chemistry/create/potential", func(writer http.ResponseWriter, request *http.Request) {
		m := model.Potentials{}
		json.NewDecoder(request.Body).Decode(&m)
		readParam := ReadAllPotentialParam(&m)
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

func (pt *PotentialController) FindPotentialElController(r *repository.PotentialsRepo, rou *mux.Router) {
	pt.Router = rou
	pt.Router.HandleFunc("/chemistry/potential", func(writer http.ResponseWriter, request *http.Request) {
		m, err := r.FindPotentialElByName(request.URL.Query().Get("name"))
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
func ReadAllPotentialParam(mod *model.Potentials) string {
	if mod.Number == "" {
		return "номеро4ек пустой"
	}
	if mod.Symbol == "" {
		return "символ пустой"
	}
	if mod.ElName == "" {
		return "нет имени элемента"
	}
	if mod.HalfReactions == "" {
		return "полуреакции нет"
	}
	if mod.LastParam == "" {
		return "нет последнего параметра"
	}
	return ""
}
