package repository

import "vitalic_project/internal/app/store"

type MainRepo struct {
	BaseVariableRepo *BaseVariableRepo
	TermodPropRepo   *TermodPropRepo
	PotentialsRepo   *PotentialsRepo
	BalanceRepo      *BalanceRepo
	InstabilityRepo  *InstabilityRepo
}

func NewMainRepo(store *store.Store) *MainRepo {
	return &MainRepo{
		BaseVariableRepo: NewBaseVariableRepo(store),
		TermodPropRepo:   NewTermodPropRepo(store),
		PotentialsRepo:   NewPotentialsRepo(store),
		BalanceRepo:      NewBalanceRepo(store),
		InstabilityRepo:  NewInstabilityRepo(store),
	}
}
