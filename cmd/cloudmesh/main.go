package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/yourusername/nebula-grid/internal/cluster"
	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/internal/db"
	"github.com/yourusername/nebula-grid/internal/server"
	"github.com/yourusername/nebula-grid/pkg/logger"
)

func main() {
	cfg := config.Load()
	logger.Init(cfg)

	gdb, err := db.Open(cfg)
	if err != nil {
		logger.Log.Fatalf("failed to open db: %v", err)
	}
	defer func() {
		sqlDB, _ := gdb.DB()
		_ = sqlDB.Close()
	}()

	// Start consensus worker (simulated)
	cons := cluster.NewConsensusWorker(gdb, cfg)
	cons.Start()
	defer cons.Stop()

	httpSrv := server.NewHTTPServer(cfg, gdb)

	// Start server in goroutine so we can gracefully shutdown
	go func() {
		logger.Log.Infof("starting server on :%s", cfg.Port)
		if err := httpSrv.ListenAndServe(); err != nil && err != httpSrv.ErrServerClosed {
			logger.Log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for signal
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	logger.Log.Info("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := httpSrv.Shutdown(ctx); err != nil {
		logger.Log.Fatalf("server forced to shutdown: %v", err)
	}

	logger.Log.Info("server exiting")
}
