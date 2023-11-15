package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

const (
	port = ":8080"
)

// crear una aplicación web con el framework Gin que tenga un endpoint /productos que responda con una lista de productos.

//Product es una estructura con los valores:Id, Nombre, Precio, Stock, Codigo, Publicado, FechaDeCreacion
type Product struct {
	Id				string	`json:"id"`
	Name    		string	`json:"name"`
	Price  			float64	`json:"price"`
	Stock   		int		`json:"stock"`
	Code 			string	`json:"code"`
	Published		bool 	`json:"published"`
	CreationDate	string	`json:"creationDate"`
}

func main() {

	// Crear un archivo productos.json con al menos seis productos (deben seguir la misma estructura ya mencionada).
	// Leer el archivo productos.json.
	fileContent, err := ioutil.ReadFile("products.json")
	if err != nil {
		log.Fatal(err)
	}

	// Deserializar el contenido del archivo en la estructura Product.
	var products []Product
	err = json.Unmarshal(fileContent, &products)
	if err != nil {
		log.Fatal(err)
	}
	
	// Imprimir la lista de productos por consola
	fmt.Println(products)

	// Imprimir la lista de productos a través del endpoint en formato JSON. El endpoint deberá ser de método GET.
	router := gin.Default()

	router.GET("/products", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H {
			"data": products,
		})
	})

	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}