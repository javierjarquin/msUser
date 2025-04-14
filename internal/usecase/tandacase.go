package usercase

import (
	"msUser/internal/domain"
	"msUser/internal/repository"
)

type TandaCase struct {
	tandaRepo repository.TandaRepo
}

func NewTandaCase(tandaRepo repository.TandaRepo) *TandaCase {
	return &TandaCase{tandaRepo: tandaRepo}
}

func (tc *TandaCase) NewTanda(tanda domain.Tanda) (domain.Tanda, error) {

	return tc.tandaRepo.NewTanda(tanda)
}

func (utc *TandaCase) UpdateTanda(tanda domain.Tanda) (domain.Tanda, error) {
	return utc.tandaRepo.UpdateTanda(tanda)
}

func (utc *TandaCase) GetTandaById(Id int) (domain.Tanda, error) {
	return utc.tandaRepo.GetTandaById(Id)
}

func (utc *TandaCase) GetTandaByUserId(Id int) ([]domain.Tanda, error) {
	return utc.tandaRepo.GetTandaByUserId(Id)
}
