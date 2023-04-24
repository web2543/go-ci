package hope

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Version struct {
	Version string `json:"version"`
}

func Getversion(c *gin.Context) {
	var version Version
	version.Version = "0.0.1"
	c.JSON(http.StatusOK, version)
}
