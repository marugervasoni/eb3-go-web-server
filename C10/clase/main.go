package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// estructura
type Product struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Quantity    int       `json:"quantity"`
	CodeValue   string    `json:"code_value"`
	IsPublished bool      `json:"is_published"`
	Expiration  time.Time `json:"expiration"`
	Price       float64   `json:"price"`
}

// estructura que simula BD en memoria
type Store struct {
	Products []Product
}

func main() {

	//carga la base de datos en memoria
	store := Store{}
	store.LoadStore()

	engine := gin.Default()

	//agrupamos en un path
	group := engine.Group("/api/v1")
	{
		//Ruta para probar el servidor
		group.GET("ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{
				"message":"pong",
			})
		})

		//subagrupación 
		//Ruta para obtener todos los productos
		prodGroup := group.Group("/products")
		{
			prodGroup.GET("/", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, gin.H{
					"data": store.Products,
				})
			})

			//implementamos path para buscar productos por precio
			prodGroup.GET("/search/:paramPrice", func(ctx *gin.Context) {
				
				//opcional
				// queryParam := ctx.Query("query")

				//recuperamos el parametro enviado en el path
				priceParam := ctx.Param("paramPrice")

				//casteamos precio porque es float y tenemos un string
				castedPrice, err := strconv.ParseFloat(priceParam, 64)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid price parameter",
					})
					return 	//p que salga de la función una vez que falla 
				}

				var result []Product
				for _, product := range store.Products {
					//si precio d´producto es mayor al precio casteado lo guaro en mi slice result
					if product.Price > castedPrice {
						result = append(result, product)
					}
				}

				ctx.JSON(http.StatusOK, gin.H{
					"data":result,
				})
			})

			// ruta /productparams que tome todos los datos de la estructura de un producto por parámetro y lo devuelva en forma de JSON
			prodGroup.GET("/productparams", func(ctx *gin.Context) {
				
				//recuperamos parametros enviados
				idParam := ctx.Query("id")
				nameParam := ctx.Query("name")
				quantityParam := ctx.Query("quantity")
				codeValueParam := ctx.Query("code_value")
				IsPublishedParam := ctx.Query("is_published")
				expirationParam := ctx.Query("expiration")
				priceParam := ctx.Query("price")

				//casteamos
				castedPrice, err := strconv.ParseFloat(priceParam, 64)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid price parameter",
					})
					return 
				}
				castedQuantity, err := strconv.Atoi(quantityParam)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid quantity parameter",
					})
					return 
				}
				castedIsPublished, err := strconv.ParseBool(IsPublishedParam)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid isPublished parameter",
					})
					return 
				}
				timeString := "2006-01-02 15:04:05"
				castedExpiration, err := time.Parse(timeString, expirationParam)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid expiration parameter",
					})
					return 
				}

				//cargamos el nuevo producto
				newProduct := Product{
					Id:          idParam,
					Name:        nameParam,
					Quantity:    castedQuantity,
					CodeValue:   codeValueParam,
					IsPublished: castedIsPublished,
					Expiration:  castedExpiration,
					Price:       castedPrice,
				}

				//Insertar este último producto a la lista de productos y verificar si lo podemos tomar con la ruta /products/:id
				store.Products = append(store.Products, newProduct)
				
				ctx.JSON(http.StatusOK, gin.H{
					"data": newProduct,
				})
			})
			
			// /searchbyquantity: devuelva una lista de productos que estén entre ciertas cantidades de stock. (ejemplo: los productos que 
			// tengan entre 300 y 400 unidades). *pasar los límites de las cantidades por parámetro
			prodGroup.GET("/searchbyquantity", func(ctx *gin.Context) {
				
				quantityParam1 := ctx.Query("quantity1")
				quantityParam2 := ctx.Query("quantity2")

				//casteamos 
				castedQuantity1, err := strconv.Atoi(quantityParam1)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid quantityParameter1",
					})
					return 
				}

				castedQuantity2, err := strconv.Atoi(quantityParam2)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid quantityParameter2",
					})
					return 
				}

				var result []Product
				for _, product := range store.Products {
					//si unidades de producto se correponden con las cantidades
					if product.Quantity >= castedQuantity1 && product.Quantity <= castedQuantity2 {
						result = append(result, product)
					}
				}

				ctx.JSON(http.StatusOK, gin.H{
					"data": result,
				})
			})
			
			// /buy -> endpoint que brinde el detalle de una compra. Por parámetro se deberá pasar el code_value del producto
			//  y la cantidad de unidades a comprar. El detalle de la compra deberá ser: nombre del producto, cantidad y precio total. 
			prodGroup.GET("/buy", func(ctx *gin.Context) {
				
				codeValueParam := ctx.Query("code_value")
				quantityToBuyParam := ctx.Query("quantity_to_buy")

				//casteamos 
				castedQuantity, err := strconv.Atoi(quantityToBuyParam)
				if err!= nil {
					ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
						"message":"ivalid quantity parameter",
					})
					return 
				}

				var result []Product
				for _, product := range store.Products {
					//si code_value de producto se correponden con la busqueda
					//comparamos strings ignorando mayusculas y/o minusculas
					if  strings.EqualFold(product.CodeValue, codeValueParam) && product.Quantity >= castedQuantity {
						result = append(result, product)
						break //solo necesitamos 1 producto coincidente
					}
				}

				if len(result) == 0 {
					ctx.JSON(http.StatusNotFound, gin.H{
						"message": "No matching products found",
					})
					return
				}
			
				totalPrice := float64(castedQuantity) * result[0].Price
				formattedTotalPrice := strconv.FormatFloat(totalPrice, 'f', 2, 64)
			
				details := fmt.Sprintf("Details: Product name: %s, quantity: %d, total price: $%s.", result[0].Name, castedQuantity, formattedTotalPrice)			
				ctx.JSON(http.StatusOK, gin.H{
					"data": details,				
				})
			})
		}
	}

	if err := engine.Run(":8080"); err != nil {
		log.Fatal(err)
	}

}

// metodo LoadStore de caraga de productos
func (s *Store) LoadStore() { 
	s.Products = []Product{
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