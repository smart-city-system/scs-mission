package models

import (
	"time"

	"github.com/google/uuid"
)

// Alarm represents an alarm in the SCS system.
type Alarm struct {
	Base
	PremiseID   uuid.UUID `json:"premise_id"`
	Premise     *Premise  `json:"premise,omitempty" gorm:"foreignKey:PremiseID"`
	Type        string    `json:"type"`
	Description string    `json:"description"`
	Severity    string    `json:"severity" gorm:"check:severity IN ('low', 'medium', 'high')"`
	TriggeredAt time.Time `json:"triggered_at" gorm:"type:timestamptz;default:CURRENT_TIMESTAMP"`
}
