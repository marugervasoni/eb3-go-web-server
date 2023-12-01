package products

import (
	"context"
	"database/sql"
	"errors"

	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
)

var(
	ErrPrepareStatement = errors.New("error prepare statement")
	ErrExecStatement = errors.New("error exec statement")
	ErrLastInsertedId = errors.New("error last inserted id")	
)

type repositorymysql struct {
	db *sql.DB
}

//NewMysqlReository
func NewMysqlRepository(db *sql.DB) Repository{
	return &repositorymysql{db: db}
}

//Create
func (r *repositorymysql) Create(ctx context.Context, product domain.Product) (domain.Product, error) {
	statement, err := r.db.Prepare(QueryInsertProduct)
	if err!= nil {
		return domain.Product{}, ErrPrepareStatement
	}

	defer statement.Close()

	result, err := statement.Exec(
		product.Name,
		product.Quantity,
		product.CodeValue,
		product.IsPublished,
		product.Expiration,
		product.Price,
	)
	if err != nil {
		return domain.Product{}, ErrExecStatement
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return domain.Product{}, ErrLastInsertedId
	}

	product.Id = int(lastId)

	return product, nil
}


//GetAll
func (r *repositorymysql) GetAll(ctx context.Context) ([]domain.Product, error){
	panic("implement me")
}


//GetByID
func (r *repositorymysql) GetByID(ctx context.Context, id string) (domain.Product, error){
	panic("implement me")
}


//Update
func (r *repositorymysql) Update(ctx context.Context, product domain.Product, id string) (domain.Product, error){
	panic("implement me")
}
