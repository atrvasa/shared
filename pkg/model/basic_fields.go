package model

import (
	"time"

	"github.com/google/uuid"
)

type BaseEntity struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`

	Reserved          bool       `gorm:"not null;default:false;index" json:"reserved"`
	ReservedSessionID *uuid.UUID `gorm:"type:uuid;index" json:"reserved_sessionId,omitempty"`
	ReservedAt        *time.Time `gorm:"index" json:"reservedAt,omitempty"`

	UpdatedSessionID *uuid.UUID `gorm:"type:uuid" json:"updated_sessionId,omitempty"`
	UpdatedAt        time.Time  `gorm:"not null;autoUpdateTime" json:"updated_at"`

	// CreatedSessionID uuid.UUID `gorm:"type:uuid;not null;index" json:"createdSessionId"`
	CreatedSessionID *uuid.UUID `gorm:"type:uuid;index" json:"created_session_id"`
	CreatedAt        time.Time  `gorm:"not null;autoCreateTime" json:"created_at"`

	Deleted          bool       `gorm:"not null;default:false;index" json:"deleted"`
	DeletedSessionID *uuid.UUID `gorm:"type:uuid;index" json:"deleted_sessionId,omitempty"`
	DeletedAt        *time.Time `gorm:"index" json:"deletedAt,omitempty"`
}

type BaseComposite struct {
	ID uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
}
