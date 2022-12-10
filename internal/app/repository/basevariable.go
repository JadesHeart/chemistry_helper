package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/store"
)

type BaseVariableRepo struct {
	store *store.Store
}

func NewBaseVariableRepo(st *store.Store) *BaseVariableRepo {
	return &BaseVariableRepo{
		store: st,
	}
}
func (r *BaseVariableRepo) Create(bv *model.BaseVariable) (*model.BaseVariable, error) {
	if model, _ := r.FindBaseVarElByName(bv.ElName); model != nil {
		return nil, errors.New("такой элемент уже есть")
	}
	if err := r.store.GetterDB().QueryRow(
		"INSERT INTO "+
			"lowercase_characteristics(el_name,electronic_configuration,stability_oxidation_state_configuration,melting_point,boiling_point,chemical_compounds)"+
			"VALUES ($1,$2,$3,$4,$5,$6) RETURNING id",
		bv.ElName,
		bv.ElectronicConfiguration,
		bv.StabilityOxidation,
		bv.MeltingPoint,
		bv.BoilingPoint,
		bv.ChemicalCompounds,
	).Scan(&bv.Id); err != nil {
		return nil, err
	}
	return bv, nil
}

func (r *BaseVariableRepo) FindBaseVarElByName(elName string) (*model.BaseVariable, error) {
	bv := &model.BaseVariable{}
	row := r.store.GetterDB().QueryRow(
		"SELECT * FROM lowercase_characteristics WHERE el_name=$1",
		elName)
	err := row.Scan(
		&bv.Id,
		&bv.ElName,
		&bv.ElectronicConfiguration,
		&bv.StabilityOxidation,
		&bv.MeltingPoint,
		&bv.BoilingPoint,
		&bv.ChemicalCompounds)
	if err == sql.ErrNoRows {
		return nil, err
	} else if err != nil {
		fmt.Println("Unexpected error: ", err.Error())
	}
	return bv, nil
}
