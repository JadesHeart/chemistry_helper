package repository

import (
	"database/sql"
	"fmt"
	"vitalic_project/internal/app/model"
	"vitalic_project/internal/app/store"
)

type PotentialsRepo struct {
	store *store.Store
}

func NewPotentialsRepo(st *store.Store) *PotentialsRepo {
	return &PotentialsRepo{
		store: st,
	}
}
func (r *PotentialsRepo) Create(pt *model.Potentials) (*model.Potentials, error) {
	if err := r.store.GetterDB().QueryRow(
		"INSERT INTO "+
			"potential(number,symbol,el_name,half_reactions,last_param)"+
			"VALUES ($1,$2,$3,$4,$5) RETURNING id",
		pt.Number,
		pt.Symbol,
		pt.ElName,
		pt.HalfReactions,
		pt.LastParam,
	).Scan(&pt.Id); err != nil {
		return nil, err
	}
	return pt, nil
}

func (r *PotentialsRepo) FindPotentialElByName(elName string) ([]*model.Potentials, error) {
	m := &model.Potentials{}
	var pt []*model.Potentials
	rows, _ := r.store.GetterDB().Query("SELECT * FROM potential WHERE el_name=$1", elName)
	for rows.Next() {
		err := rows.Scan(
			&m.Id,
			&m.Number,
			&m.Symbol,
			&m.ElName,
			&m.HalfReactions,
			&m.LastParam)
		if err == sql.ErrNoRows {
			return nil, err
		} else if err != nil {
			fmt.Println("Unexpected error: ", err.Error())
		}
		pt = append(pt, m)
		m = &model.Potentials{}
	}
	return pt, nil
}
