package db_test

import (
	"testing"

	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/internal/db"
	"github.com/yourusername/nebula-grid/internal/models"
)

func TestOpenSqliteInMemory(t *testing.T) {
	cfg := &config.Config{DBDriver: "sqlite"}
	gdb, err := db.Open(cfg)
	if err != nil {
		t.Fatalf("failed to open db: %v", err)
	}

	// create node
	n := &models.Node{Name: "test-node", Address: "127.0.0.1:9000"}
	if err := gdb.Create(n).Error; err != nil {
		t.Fatalf("create node failed: %v", err)
	}

	var out models.Node
	if err := gdb.First(&out, "name = ?", "test-node").Error; err != nil {
		t.Fatalf("query node failed: %v", err)
	}
	if out.Address != n.Address {
		t.Fatalf("unexpected address: %s", out.Address)
	}
}
