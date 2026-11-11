package cluster

import (
	"context"
	"sync"
	"time"

	"github.com/yourusername/nebula-grid/internal/config"
	"github.com/yourusername/nebula-grid/internal/models"
	"github.com/yourusername/nebula-grid/pkg/logger"
	"gorm.io/gorm"
)

// ConsensusWorker is a simulated leader election worker. It periodically selects one node to be leader.
type ConsensusWorker struct {
	db       *gorm.DB
	cfg      *config.Config
	cancel   context.CancelFunc
	wg       sync.WaitGroup
	stopped  bool
	stopLock sync.Mutex
}

func NewConsensusWorker(db *gorm.DB, cfg *config.Config) *ConsensusWorker {
	return &ConsensusWorker{db: db, cfg: cfg}
}

func (cw *ConsensusWorker) Start() {
	ctx, cancel := context.WithCancel(context.Background())
	cw.cancel = cancel
	cw.wg.Add(1)
	go cw.run(ctx)
}

func (cw *ConsensusWorker) Stop() {
	cw.stopLock.Lock()
	if cw.stopped {
		cw.stopLock.Unlock()
		return
	}
	cw.stopped = true
	cw.stopLock.Unlock()
	if cw.cancel != nil {
		cw.cancel()
	}
	cw.wg.Wait()
}

func (cw *ConsensusWorker) run(ctx context.Context) {
	defer cw.wg.Done()
	ticker := time.NewTicker(cw.cfg.ConsensusInterval)
	defer ticker.Stop()

	// Run initial election
	cw.elect()

	for {
		select {
		case <-ctx.Done():
			logger.Log.Info("consensus worker stopped")
			return
		case <-ticker.C:
			cw.elect()
		}
	}
}

// elect selects a leader deterministically for demonstration: the oldest created node (lowest CreatedAt) becomes leader.
func (cw *ConsensusWorker) elect() {
	var nodes []models.Node
	if err := cw.db.Order("created_at asc").Find(&nodes).Error; err != nil {
		logger.Log.Errorf("consensus: failed to list nodes: %v", err)
		return
	}
	if len(nodes) == 0 {
		return
	}
	leaderID := nodes[0].ID

	// Update nodes: set IsLeader flag only for the chosen one
	if err := cw.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&models.Node{}).Where("id <> ?", leaderID).Update("is_leader", false).Error; err != nil {
			return err
		}
		if err := tx.Model(&models.Node{}).Where("id = ?", leaderID).Update("is_leader", true).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		logger.Log.Errorf("consensus: failed to update leader: %v", err)
		return
	}
	logger.Log.Infof("consensus: elected leader node id=%d name=%s", leaderID, nodes[0].Name)
}
