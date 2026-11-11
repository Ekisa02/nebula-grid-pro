package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/nebula-grid/internal/api"
	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/pkg/logger"
	"gorm.io/gorm"
)

// NewRouter constructs the Gin engine with routes and middleware.
func NewRouter(db *gorm.DB) *gin.Engine {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	h := api.NewHandler(db)
	v1 := r.Group("/api/v1")
	{
		v1.POST("/nodes", h.RegisterNode)
		v1.GET("/nodes", h.ListNodes)
		v1.POST("/services", h.RegisterService)
		v1.GET("/services", h.ListServices)
	}

	// health
	r.GET("/healthz", func(c *gin.Context) { c.Status(http.StatusOK) })

	return r
}

// NewHTTPServer returns an http.Server ready to ListenAndServe.
func NewHTTPServer(cfg *config.Config, db *gorm.DB) *http.Server {
	r := NewRouter(db)
	return &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: r,
	}
}
