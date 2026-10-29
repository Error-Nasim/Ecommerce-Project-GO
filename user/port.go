package user

import (
	"ecommerece/domain"
	userHandler "ecommerece/rest/handlers/user"
)

type Service interface {
	userHandler.Service
}

type UserRepo interface {
	Create(p domain.User) (*domain.User, error)
	Find(email, pass string) (*domain.User, error)
}
