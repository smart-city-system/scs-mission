package models

import "github.com/google/uuid"

// GuidanceStep represents an individual step in a guidance template
// @Description Individual step within a guidance template procedure
type GuidanceStep struct {
	Base
	GuidanceTemplateID uuid.UUID         `json:"guidance_template_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	GuidanceTemplate   *GuidanceTemplate `json:"guidance_template,omitempty" gorm:"foreignKey:GuidanceTemplateID"`
	StepNumber         int               `json:"step_number" example:"1"`
	Title              string            `json:"title" example:"Assess the situation"`
	Description        string            `json:"description" example:"Quickly evaluate the severity and scope of the fire"`
}
