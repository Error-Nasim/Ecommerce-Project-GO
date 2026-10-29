package product

import "ecommerece/domain"

type service struct {
	prdRepo ProductRepo
}

func NewService(prdRepo ProductRepo) Service {
	return &service{
		prdRepo: prdRepo,
	}
}

func (svc *service) Create(prdct domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Create(prdct)
}

func (svc *service) Get(id int) (*domain.Product, error) {
	return svc.prdRepo.Get(id)
}

func (svc *service) List() ([]*domain.Product, error) {
	return svc.prdRepo.List()
}

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(prdct)
}

func (svc *service) Delete(id int) error {
	return svc.prdRepo.Delete(id)
}
