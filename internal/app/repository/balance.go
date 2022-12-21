package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/store"
)

type BalanceRepo struct {
	store *store.Store
}

func NewBalanceRepo(st *store.Store) *BalanceRepo {
	return &BalanceRepo{
		store: st,
	}
}
func (r *BalanceRepo) Create(bl *model.BalanceConst) (*model.BalanceConst, error) {
	if model, _ := r.FindBalanceElByName(bl.ElName); model != nil {
		return nil, errors.New("такой элемент уже есть")
	}
	if err := r.store.GetterDB().QueryRow(
		"INSERT INTO "+
			"balance(el_name,formula,first_param,second_param)"+
			"VALUES ($1,$2,$3,$4) RETURNING id",
		bl.ElName,
		bl.Formula,
		bl.FirstParam,
		bl.SecondParam,
	).Scan(&bl.Id); err != nil {
		return nil, err
	}
	return bl, nil
}

func (r *BalanceRepo) FindBalanceElByName(elName string) (*model.BalanceConst, error) {
	bl := &model.BalanceConst{}
	row := r.store.GetterDB().QueryRow(
		"SELECT * FROM balance WHERE el_name=$1",
		elName)
	err := row.Scan(
		&bl.Id,
		&bl.ElName,
		&bl.Formula,
		&bl.FirstParam,
		&bl.SecondParam,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		fmt.Println("Unexpected error: ", err.Error())
	}
	return bl, nil
}

func (r *BalanceRepo) FindBalanceElByFormula(formula string) (*model.BalanceConst, error) {
	bl := &model.BalanceConst{}
	row := r.store.GetterDB().QueryRow(
		"SELECT * FROM balance WHERE formula=$1",
		formula)
	err := row.Scan(
		&bl.Id,
		&bl.ElName,
		&bl.Formula,
		&bl.FirstParam,
		&bl.SecondParam,
	)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		fmt.Println("Unexpected error: ", err.Error())
	}
	return bl, nil
}
