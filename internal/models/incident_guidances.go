package models

import "github.com/google/uuid"

// IncidentGuidance represents guidance assigned to a specific incident
// @Description Guidance assignment linking an incident to a guidance template with assignee information
type IncidentGuidance struct {
	Base
	IncidentID            *uuid.UUID             `json:"incident_id" gorm:"uniqueIndex:idx_incident_guidance" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Incident              *Incident              `json:"incident,omitempty" gorm:"foreignKey:IncidentID"`
	GuidanceTemplateID    *uuid.UUID             `json:"guidance_template_id" gorm:"uniqueIndex:idx_incident_guidance" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	GuidanceTemplate      *GuidanceTemplate      `json:"guidance_template,omitempty" gorm:"foreignKey:GuidanceTemplateID"`
	AssignerID            *uuid.UUID             `json:"assigner_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Assigner              *User                  `json:"assigner,omitempty" gorm:"foreignKey:AssignerID"`
	AssigneeID            *uuid.UUID             `json:"assignee_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Assignee              *User                  `json:"assignee,omitempty" gorm:"foreignKey:AssigneeID"`
	IncidentGuidanceSteps []IncidentGuidanceStep `json:"incident_guidance_steps" gorm:"foreignKey:IncidentGuidanceID"`
}
