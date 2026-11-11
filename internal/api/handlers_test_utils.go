package api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// RouterForTest builds a small router for handler unit tests.
func (h *Handler) RouterForTest() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.POST("/nodes", h.RegisterNode)
	v1.GET("/nodes", h.ListNodes)
	v1.POST("/services", h.RegisterService)
	v1.GET("/services", h.ListServices)
	return r
}
