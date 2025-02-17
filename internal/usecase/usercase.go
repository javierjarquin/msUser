package usercase

import (
	"errors"
	"msUser/internal/domain"
	"msUser/internal/repository"
	"time"
)

type UserCase struct {
	userRepo repository.UserRepo
}

func NewUserCase(userRepo repository.UserRepo) *UserCase {
	return &UserCase{userRepo: userRepo}
}

func (uc *UserCase) NewUser(user domain.User) (domain.User, error) {
	return uc.userRepo.NewUser(user)
}

func (usp *UserCase) UpdateUser(user domain.User) (domain.User, error) {
	return usp.userRepo.UpdateUser(user)
}

func (usp *UserCase) GetUserById(Id int) (domain.User, error) {
	return usp.userRepo.GetUserById(Id)
}

func (usp *UserCase) LoginUser(email string, pass string, IPAddress string, coment string) (domain.Session, error) {
	//
	// Creamos funcioalidad para loguear usuario
	//

	user, err := usp.userRepo.GetUserByEmail(email)
	if err != nil {
		return domain.Session{}, errors.New("Usuario no encontrado")
	}

	if user.Pass != pass {
		return domain.Session{}, errors.New("Contraseña incorrecta")
	}

	ses, err := usp.userRepo.NewSession(domain.Session{
		UserID:       user.ID,
		CreationDate: time.Now(),
		IPAddres:     IPAddress,
		Comments:     coment,
	})

	return domain.Session{ID: ses.ID}, err
}
