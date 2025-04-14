package usercase

import (
	"msUser/internal/domain"
	"msUser/internal/repository"
)

type TandaPagoCase struct {
	tandapagoRepo repository.TandaPagoRepo
}

func NewTandaPagoCase(tandapagoRepo repository.TandaPagoRepo) *TandaPagoCase {
	return &TandaPagoCase{tandapagoRepo: tandapagoRepo}
}

func (tc *TandaPagoCase) NewTandaPago(tandapago domain.TandaPago) (domain.TandaPago, error) {

	return tc.tandapagoRepo.NewTandaPago(tandapago)
}

func (utc *TandaPagoCase) UpdateTandaPago(tandapago domain.TandaPago) (domain.TandaPago, error) {
	return utc.tandapagoRepo.UpdateTandaPago(tandapago)
}

func (utc *TandaPagoCase) GetTandaPagoByTandaUsuarioId(Id int) ([]domain.TandaPago, error) {
	return utc.tandapagoRepo.GetTandaPagoByTandaUsuarioId(Id)
}
