package models

import (
	"time"

	"github.com/google/uuid"
)

// Base represents the common fields for all models
// @Description Common base fields for all entities
type Base struct {
	ID        uuid.UUID `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime" example:"2023-01-01T00:00:00Z"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime" example:"2023-01-01T00:00:00Z"`
}
