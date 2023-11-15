package main

import "github.com/gin-gonic/gin"

func main() {

	//creo router con gin
	router := gin.Default()

	//capturo solicitud GET "/hello-world"
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})

	//corremos nustro servidor sobre puerto 8080
	router.Run()
}