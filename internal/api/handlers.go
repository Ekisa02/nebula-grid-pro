package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/nebula-grid/internal/models"
	"gorm.io/gorm"
)

// Handler groups dependencies for HTTP handlers.
type Handler struct {
	db *gorm.DB
}

func NewHandler(db *gorm.DB) *Handler {
	return &Handler{db: db}
}

// RegisterNode registers a new node.
func (h *Handler) RegisterNode(c *gin.Context) {
	var n models.Node
	if err := c.ShouldBindJSON(&n); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&n).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, n)
}

// ListNodes returns all nodes.
func (h *Handler) ListNodes(c *gin.Context) {
	var nodes []models.Node
	if err := h.db.Find(&nodes).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, nodes)
}

// RegisterService registers a service record.
func (h *Handler) RegisterService(c *gin.Context) {
	var s models.Service
	if err := c.ShouldBindJSON(&s); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.db.Create(&s).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, s)
}

// ListServices returns all services.
func (h *Handler) ListServices(c *gin.Context) {
	var services []models.Service
	if err := h.db.Find(&services).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, services)
}
