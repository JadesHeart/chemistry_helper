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

type BalanceController struct {
	Router *mux.Router
}

func NewBalanceController() *BalanceController {
	return &BalanceController{}
}
func (blc *BalanceController) CreateBalanceController(r *repository.BalanceRepo, rou *mux.Router) {
	blc.Router = rou
	blc.Router.HandleFunc("/chemistry/create/balance", func(writer http.ResponseWriter, request *http.Request) {
		m := model.BalanceConst{}
		json.NewDecoder(request.Body).Decode(&m)
		readParam := ReadAllBalanceParam(&m)
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

func (blc *BalanceController) FindBalanceElController(r *repository.BalanceRepo, rou *mux.Router) {
	blc.Router = rou
	blc.Router.HandleFunc("/chemistry/balance", func(writer http.ResponseWriter, request *http.Request) {
		query := request.URL.Query().Get("name")
		m, err := r.FindBalanceElByName(request.URL.Query().Get("name"))
		if err == sql.ErrNoRows {
		} else if m != nil && err == nil {
			returnedModel, err := json.Marshal(m)
			if err != nil {
				fmt.Fprintf(writer, "Неудалось распокавать json")
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(returnedModel)
		}
		m1, err1 := r.FindBalanceElByFormula(query)
		if err1 == sql.ErrNoRows && err == sql.ErrNoRows {
			fmt.Fprintf(writer, "неверное имя элемента")
		} else if m1 != nil && err1 == nil {
			returnedModel, err := json.Marshal(m1)
			if err != nil {
				fmt.Fprintf(writer, "Неудалось распокавать json")
			}
			writer.Header().Set("Content-Type", "application/json")
			writer.WriteHeader(http.StatusOK)
			writer.Write(returnedModel)
		}
	})
}
func ReadAllBalanceParam(mod *model.BalanceConst) string {
	if mod.ElName == "" {
		return "нет имени элемента или оно пустое"
	}
	if mod.Formula == "" {
		return "формула пустая"
	}
	if mod.FirstParam == "" {
		return "нет первого параметра"
	}
	if mod.SecondParam == "" {
		return "нет второго параметра"
	}
	return ""
}
