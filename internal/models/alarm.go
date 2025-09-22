package models

import (
	"time"

	"github.com/google/uuid"
)

// Alarm represents an alarm in the SCS system
// @Description Alarm entity triggered by security events
type Alarm struct {
	Base
	PremiseID   uuid.UUID `json:"premise_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Premise     *Premise  `json:"premise,omitempty" gorm:"foreignKey:PremiseID"`
	Type        string    `json:"type" example:"fire"`
	Description string    `json:"description" example:"Fire alarm triggered in main building"`
	Severity    string    `json:"severity" gorm:"check:severity IN ('low', 'medium', 'high')" example:"high" enums:"low,medium,high"`
	TriggeredAt time.Time `json:"triggered_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP" example:"2023-01-01T00:00:00Z"`
}
