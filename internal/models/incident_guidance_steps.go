package models

import (
	"time"

	"github.com/google/uuid"
)

// IncidentGuidanceStep represents a step in an incident guidance procedure
// @Description Individual step within an incident guidance with completion tracking
type IncidentGuidanceStep struct {
	Base
	IncidentGuidanceID uuid.UUID         `json:"incident_guidance_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	IncidentGuidance   *IncidentGuidance `json:"incident_guidance,omitempty" gorm:"foreignKey:IncidentGuidanceID"`
	StepNumber         int64             `json:"step_number" example:"1"`
	Title              string            `json:"title" example:"Assess the situation"`
	Description        string            `json:"description" example:"Quickly evaluate the severity and scope of the incident"`
	IsCompleted        bool              `json:"is_completed" gorm:"default:false" example:"false"`
	CompletedAt        *time.Time        `json:"completed_at,omitempty" example:"2023-01-01T00:00:00Z"`
}
