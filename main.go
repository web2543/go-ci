package main

import (
	"argo/api/hope"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/version",
		hope.Getversion)
	return r

}

func main() {
	r := setupRouter()
	r.Run(":3001")
}
