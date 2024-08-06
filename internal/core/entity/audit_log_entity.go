package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/datatypes"
)

type AuditLog struct {
	ID         uuid.UUID      `json:"id" gorm:"primaryKey;type:uuid;default:uuid_generate_v4();"`
	UserId     *uuid.UUID     `json:"user_id" gorm:"type:uuid"`
	EntityType string         `json:"entity_type" gorm:"not null"`
	EntityId   uuid.UUID      `json:"entity_id" gorm:"type:uuid"`
	Action     string         `json:"action"`
	Data       datatypes.JSON `json:"data" gorm:"type:jsonb"`
	CreatedAt  time.Time      `json:"created_at"`

	// relationships
	// User User `json:"user" gorm:"references:ID"`
}

func (AuditLog) TableName() string {
	return "audit_log"
}
