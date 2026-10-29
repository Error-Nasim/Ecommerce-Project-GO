package user

import (
	"ecommerece/domain"
)

type service struct {
	usrRepo UserRepo
}

func NewService(usrRepo UserRepo) Service {
	return &service{
		usrRepo: usrRepo,
	}
}

func (svc *service) Create(user domain.User) (*domain.User, error) {
	usr, err := svc.usrRepo.Create(user)
	if err != nil || usr == nil {
		return nil, err
	}
	return usr, nil
}
func (svc *service) Find(email string, password string) (*domain.User, error) {
	usr, err := svc.usrRepo.Find(email, password)
	if err != nil || usr == nil {
		return nil, err
	}
	return usr, nil
}
