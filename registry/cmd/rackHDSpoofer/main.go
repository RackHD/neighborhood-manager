package main

import (
	"log"
	"net/http"

	"gopkg.in/gin-gonic/gin.v1"
)

func main() {
	r := gin.Default()
	r.GET("/api/2.0/nodes", func(c *gin.Context) {
		c.String(http.StatusOK, "")
	})
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatalln(err)
	} // listen and serve on addr:port
}
