package repository

import "vitalic_project/internal/app/store"

type MainRepo struct {
	BaseVariableRepo *BaseVariableRepo
	TermodPropRepo   *TermodPropRepo
}

func NewMainRepo(store *store.Store) *MainRepo {
	return &MainRepo{
		BaseVariableRepo: NewBaseVariableRepo(store),
		TermodPropRepo:   NewTermodPropRepo(store),
	}
}
