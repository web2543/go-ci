package main

import (
	"argo/api/hope"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	group := r.Group("/go")
	group.GET("/version",
		hope.Getversion)
	return r

}

func main() {
	r := setupRouter()
	r.Run(":3000")
}
