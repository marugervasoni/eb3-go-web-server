package products

import (
	"context"
	"errors"
	"log"

	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
)

var (
	ErrEmpty = errors.New("empty list")
	ErrNotFound = errors.New("products not found")
)

type repository struct {
	db []domain.Product
}


//NewMemoryReository
func NewMemoryRepository(db []domain.Product) Repository{
	return &repository{db: db}
}


//GetAll
func (r *repository) GetAll(ctx context.Context) ([]domain.Product, error){

	//--implementación de AddValueToContext
	contentContext := ctx.Value("rol")

	if contentContext != ""{
		log.Println("El contenido del contexto es:", contentContext)
	}
	//--

	//si len de db es menor, la bd esta vacia, retornamos error
	if len(r.db) < 1 {
		return []domain.Product{}, ErrEmpty
	}
	//si no hay error retorno el slice y nil para el error
	return r.db, nil
}


//GetByID
func (r *repository) GetByID(ctx context.Context, id string) (domain.Product, error){
	var result domain.Product
	for _, value := range r.db {
		if value.Id == id {
			result = value
			break
		}
	}
	if result.Id == "" {
		return domain.Product{}, ErrNotFound
	}
	return result, nil
}


//Update
func (r *repository) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error){
	var result domain.Product
	for key, value := range r.db {
		if value.Id == id {
			product.Id = id
			r.db[key] = product
			result = r.db[key]
			break
		}
	}
	if result.Id == "" {
		return domain.Product{}, ErrNotFound
	}
	return result, nil
}