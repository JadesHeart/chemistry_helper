package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/store"
)

type TermodPropRepo struct {
	store *store.Store
}

func NewTermodPropRepo(st *store.Store) *TermodPropRepo {
	return &TermodPropRepo{
		store: st,
	}
}
func (r *TermodPropRepo) Create(tr *model.TermodProp) (*model.TermodProp, error) {
	if model, _ := r.FindTherByName(tr.ElName); model != nil {
		//fmt.Println(model)
		return nil, errors.New("я нашёл его")
	}
	fmt.Println(r.FindTherByName(tr.ElName))
	fmt.Println("Жопаааааааааааа")
	if err := r.store.GetterDB().QueryRow(
		"INSERT INTO "+
			"thermodynamic_characteristics(el_name,first_param,second_param,third_param,fourth_param)"+
			"VALUES ($1,$2,$3,$4,$5) RETURNING id",
		tr.ElName,
		tr.FirstParam,
		tr.SecondParam,
		tr.ThirdParam,
		tr.FourthParam,
	).Scan(&tr.Id); err != nil {
		fmt.Println("Так ну всё хуёво оно не добавилось")
		return nil, err
	}
	return tr, nil
}

func (r *TermodPropRepo) FindTherByName(elName string) (*model.TermodProp, error) {
	tr := &model.TermodProp{}
	row := r.store.GetterDB().QueryRow(
		"SELECT * FROM thermodynamic_characteristics WHERE el_name=$1",
		elName)
	err := row.Scan(
		&tr.Id,
		&tr.ElName,
		&tr.FirstParam,
		&tr.SecondParam,
		&tr.ThirdParam,
		&tr.FourthParam,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		fmt.Println("Unexpected error: ", err.Error())
	}
	return tr, nil
}
