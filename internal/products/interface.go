package products

import (
	"context"
	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
)

//pasamos a tener la interfaz en un solo archivo ya que es comun a mis repositorys
type Repository interface {
	// Create(ctx context.Context, product domain.Product) ([]domain.Product, error)
	GetAll(ctx context.Context) ([]domain.Product, error)
	GetByID(ctx context.Context, id string) (domain.Product, error)
	Update(ctx context.Context, product domain.Product, id string) (domain.Product, error)
	// Delete(ctx context.Context, id int) error
}