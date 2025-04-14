package usercase

import (
	"msUser/internal/domain"
	"msUser/internal/repository"
)

type TandaUsuarioCase struct {
	tandausuarioRepo repository.TandaUsuarioRepo
}

func NewTandaUsuarioCase(tandausuarioRepo repository.TandaUsuarioRepo) *TandaUsuarioCase {
	return &TandaUsuarioCase{tandausuarioRepo: tandausuarioRepo}
}

func (tc *TandaUsuarioCase) NewTandaUsuario(tandausuario domain.TandaUsuario) (domain.TandaUsuario, error) {

	return tc.tandausuarioRepo.NewTandaUsuario(tandausuario)
}

func (utc *TandaUsuarioCase) UpdateTandaUsuario(tandausuario domain.TandaUsuario) (domain.TandaUsuario, error) {
	return utc.tandausuarioRepo.UpdateTandaUsuario(tandausuario)
}

func (utc *TandaUsuarioCase) GetTandaUsuarioByTandaId(Id int) ([]domain.TandaUsuario, error) {
	return utc.tandausuarioRepo.GetTandaUsuarioByTandaId(Id)
}
