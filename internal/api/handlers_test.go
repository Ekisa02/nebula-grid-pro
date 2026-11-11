package api_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/yourusername/nebula-grid/internal/api"
	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/internal/db"
)

func setupRouter(t *testing.T) *gin.Engine {
	cfg := &config.Config{DBDriver: "sqlite"}
	gdb, err := db.Open(cfg)
	if err != nil {
		t.Fatalf("db open failed: %v", err)
	}
	return api.NewHandler(gdb).RouterForTest()
}

// To keep handler tests simple, extend Handler with a small util used only in tests.
// We recreate minimal behavior here by using the handler to register routes.

func (h *api.Handler) RouterForTest() *gin.Engine {
	r := gin.New()
	v1 := r.Group("/api/v1")
	v1.POST("/nodes", h.RegisterNode)
	v1.GET("/nodes", h.ListNodes)
	v1.POST("/services", h.RegisterService)
	v1.GET("/services", h.ListServices)
	return r
}

func TestRegisterAndListNode(t *testing.T) {
	cfg := &config.Config{DBDriver: "sqlite"}
	gdb, err := db.Open(cfg)
	if err != nil {
		t.Fatalf("db open failed: %v", err)
	}
	h := api.NewHandler(gdb)
	r := h.RouterForTest()

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/api/v1/nodes", strings.NewReader(`{"name":"n1","address":"127.0.0.1:9001"}`))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)
	if w.Code != http.StatusCreated {
		t.Fatalf("expected 201 created got %d body=%s", w.Code, w.Body.String())
	}

	// list
	w = httptest.NewRecorder()
	req = httptest.NewRequest("GET", "/api/v1/nodes", nil)
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Fatalf("expected 200 got %d body=%s", w.Code, w.Body.String())
	}
}
