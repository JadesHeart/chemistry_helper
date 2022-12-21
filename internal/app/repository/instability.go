package repository

import (
	"database/sql"
	"fmt"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/store"
)

type InstabilityRepo struct {
	store *store.Store
}

func NewInstabilityRepo(st *store.Store) *InstabilityRepo {
	return &InstabilityRepo{
		store: st,
	}
}
func (r *InstabilityRepo) CreateEl(in *model.Instability) (*model.Instability, error) {
	if err := r.store.GetterDB().QueryRow(
		"INSERT INTO "+
			"instability(el_name,ligand,complex,lastParam)"+
			"VALUES ($1,$2,$3,$4) RETURNING id",
		in.ElName,
		in.Ligand,
		in.Complex,
		in.LastParam,
	).Scan(&in.Id); err != nil {
		return nil, err
	}
	return in, nil
}

func (r *InstabilityRepo) FindInstabilityElByName(elName string) ([]*model.Instability, error) {
	m := &model.Instability{}
	var pt []*model.Instability
	rows, _ := r.store.GetterDB().Query("SELECT * FROM instability WHERE el_name=$1", elName)
	for rows.Next() {
		err := rows.Scan(
			&m.Id,
			&m.ElName,
			&m.Ligand,
			&m.Complex,
			&m.LastParam,
		)
		if err == sql.ErrNoRows {
			return nil, err
		} else if err != nil {
			fmt.Println("Unexpected error: ", err.Error())
		}
		pt = append(pt, m)
		m = &model.Instability{}
	}
	return pt, nil
}
