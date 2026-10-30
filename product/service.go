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

func (svc *service) List(page, limit int64) ([]*domain.Product, error) {
	return svc.prdRepo.List(page, limit)
}

func (svc *service) Count() (int64, error) {
	return svc.prdRepo.Count()
}

func (svc *service) Update(prdct domain.Product) (*domain.Product, error) {
	return svc.prdRepo.Update(prdct)
}

func (svc *service) Delete(id int) error {
	return svc.prdRepo.Delete(id)
}
