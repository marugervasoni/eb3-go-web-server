package products

import (
	"context"
	"log"
	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
)

type Service interface {
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id string) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id string) (domain.Product, error)
}

type service struct {
	repository Repository
}


//NewServiceProduct
func NewServiceProduct(repository Repository) Service {
	return &service{repository: repository} 
}


//GetAll
func (s *service) GetAll(ctx context.Context) ([]domain.Product, error) {
	productsList, err := s.repository.GetAll(ctx)
	if err != nil {
		log.Println("[ProductsService] [GetAll] error getting all products", err)
		return []domain.Product{}, err
	}
	return productsList, nil
}


//GetByID
func (s *service) GetByID(ctx context.Context, id string) (domain.Product, error) {
	product, err := s.repository.GetByID(ctx, id)
	if err != nil {
		log.Println("[ProductsService] [GetByID] error getting product by Id", err)
		return domain.Product{}, err
	}
	return product, nil
}


//Updaate
func (s *service) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error){
	product, err := s.repository.Update(ctx, product, id)
	if err != nil {
		log.Println("[ProductsService] [Update] error getting product by Id", err)
		return domain.Product{}, err
	}
	return product, nil
}