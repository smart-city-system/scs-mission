package models

import "github.com/google/uuid"

type IncidentGuidance struct {
	Base
	IncidentID            *uuid.UUID             `json:"incident_id"  gorm:"uniqueIndex:idx_incident_guidance"`
	Incident              *Incident              `json:"incident,omitempty" gorm:"foreignKey:IncidentID"`
	GuidanceTemplateID    *uuid.UUID             `json:"guidance_template_id" gorm:"uniqueIndex:idx_incident_guidance"`
	GuidanceTemplate      *GuidanceTemplate      `json:"guidance_template,omitempty" gorm:"foreignKey:GuidanceTemplateID"`
	AssignerID            *uuid.UUID             `json:"assigner_id"`
	Assigner              *User                  `json:"assigner,omitempty" gorm:"foreignKey:AssignerID"`
	AssigneeID            *uuid.UUID             `json:"assignee_id"`
	Assignee              *User                  `json:"assignee,omitempty" gorm:"foreignKey:AssigneeID"`
	IncidentGuidanceSteps []IncidentGuidanceStep `json:"incident_guidance_steps" gorm:"foreignKey:IncidentGuidanceID"`
}
