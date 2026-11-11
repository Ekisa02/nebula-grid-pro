package models

import (
	"time"

	"gorm.io/gorm"
)

// Node represents a registered node in the distributed system.
type Node struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name     string `gorm:"uniqueIndex;size:128" json:"name" binding:"required"`
	Address  string `gorm:"size:256" json:"address" binding:"required"`
	IsLeader bool   `gorm:"default:false" json:"is_leader"`
}

// Service represents a logical service running inside the cluster.
type Service struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name    string `gorm:"uniqueIndex;size:128" json:"name" binding:"required"`
	Version string `gorm:"size:64" json:"version"`
	Owner   string `gorm:"size:128" json:"owner"`
}
