package models

import (
	"github.com/google/uuid"
)

// Incident represents a security incident in the system
// @Description Security incident entity with status tracking and guidance
type Incident struct {
	Base
	Name             string            `json:"name" example:"Fire in Building A"`
	Description      string            `json:"description" example:"Fire detected on the 3rd floor of Building A"`
	AlarmID          uuid.UUID         `json:"alarm_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Alarm            *Alarm            `json:"alarm,omitempty" gorm:"foreignKey:AlarmID"`
	Status           string            `json:"status" gorm:"check:status IN ('new', 'in_progress', 'resolved')" example:"new" enums:"new,in_progress,resolved"`
	Severity         string            `json:"severity" gorm:"check:severity IN ('low', 'medium', 'high')" example:"high" enums:"low,medium,high"`
	Location         string            `json:"location" example:"Building A, Floor 3"`
	IncidentGuidance *IncidentGuidance `json:"incident_guidance,omitempty" gorm:"foreignKey:IncidentID"`
}
