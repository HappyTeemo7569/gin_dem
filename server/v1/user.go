package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type ResGroup struct {
	Data string
	Path string
}

func other(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{Data: "other", Path: c.Request.URL.Path})
}

func detail(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"detail", c.Request.URL.Path})
}

func abstract(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"abstract", c.Request.URL.Path})

}

func login(c *gin.Context) {
	c.JSON(http.StatusOK, ResGroup{"login", c.Request.URL.Path})
}
