package products

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/marugervasoni/eb3-go-web-server/internal/domain"
	"github.com/marugervasoni/eb3-go-web-server/internal/products"
)

type Controller struct {
	service products.Service
}

func NewControllerProducts(service products.Service) *Controller{
	return &Controller{service: service}
}


func (c *Controller) HandlerGetAll() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		//validamos token (Implementar manualmente en todas las funciones por el momento)
		tokenHeader := ctx.GetHeader("tokenPostman")
		tokenEnv := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		//ejemplo de implementación de AddValueToContext (y en repository [GetAll])
		newContext := addValueToContext(ctx)
		ListProducts, err := c.service.GetAll(newContext)
		// ListProducts, err := c.service.GetAll(ctx)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
				"message": "Internal server error",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": ListProducts,
		})
	}
}


func (c *Controller) HandlerGetByID() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		//validate token
		tokenHeader := ctx.GetHeader("tokenPostman")
		tokenEnv := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		//recuperamos el id de la request
		idParam := ctx.Param("id")

		//llamamos al servicio
		product, err := c.service.GetByID(ctx, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
				"message": "Internal server error",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}


func (c *Controller) HandlerUpdate() gin.HandlerFunc{
	return func(ctx *gin.Context) {

		//validate token
		tokenHeader := ctx.GetHeader("tokenPostman")
		tokenEnv := os.Getenv("TOKEN")

		if tokenHeader == "" || tokenHeader != tokenEnv {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "invalid token",
			})
			return
		}

		//recuperamos el id de la request
		idParam := ctx.Param("id")

		var productRequest domain.Product

		err := ctx.Bind(&productRequest)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H {
				"message": "bad request",
			})
		}

		//llamamos al servicio
		product, err := c.service.Update(ctx, productRequest, idParam)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H {
				"message": "Internal server error",
			})
		}
		ctx.JSON(http.StatusOK, gin.H{
			"data": product,
		})
	}
}


//addValueToContext agrega un valor al contexto
func addValueToContext(ctx context.Context) context.Context {
	newContext := context.WithValue(ctx, "rol", "admin")
	return newContext
}