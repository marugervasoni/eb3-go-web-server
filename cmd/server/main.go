package main

import (
	"log"
	"time"
	"github.com/gin-gonic/gin"
	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
	"github.com/marugervasoni/eb3-go-web-server/internal/products"
	handlerProduct "github.com/marugervasoni/eb3-go-web-server/cmd/server/handler/products"
	handlerPing "github.com/marugervasoni/eb3-go-web-server/cmd/server/handler/ping"
)

func main() {

	//implementamos repository
	db := LoadStore()

	//Ping.
	controllerPing := handlerPing.NewControllerPing()

	//Products.
	repository := products.NewMemoryRepository(db)
	service := products.NewServiceProduct(repository)
	controllerProduct := handlerProduct.NewControllerProducts(service)

	engine := gin.Default()

	//agrupamos en un path
	group := engine.Group("/api/v1")
	{
		//Ruta para probar el servidor
		group.GET("ping", controllerPing.HandlerPing())

		//subagrupación 
		//Ruta para obtener todos los productos
		prodGroup := group.Group("/products")
		{
			//utilizamos la nueva implementación
			prodGroup.GET("/", controllerProduct.HandlerGetAll())
			prodGroup.GET("/:id", controllerProduct.HandlerGetByID())
			prodGroup.PUT("/:id", controllerProduct.HandlerUpdate())
		}
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}


// metodo LoadStore de caraga de productos
func LoadStore() []domain.Product { 
	return []domain.Product{
		{
			Id:          "1",
			Name:        "Coco Cola",
			Quantity:    350,
			CodeValue:   "CC1",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       10.5,
		},
		{
			Id:          "2",
			Name:        "Pepsito",
			Quantity:    100,
			CodeValue:   "P1",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       8.5,
		},
		{
			Id:          "3",
			Name:        "Fantastica",
			Quantity:    500,
			CodeValue:   "F1",
			IsPublished: true,
			Expiration:  time.Now(),
			Price:       5.5,
		},		
	}
}