package models

import "github.com/google/uuid"

type GuidanceStep struct {
	Base
	GuidanceTemplateID uuid.UUID         `json:"guidance_template_id"`
	GuidanceTemplate   *GuidanceTemplate `json:"guidance_template,omitempty" gorm:"foreignKey:GuidanceTemplateID"`
	StepNumber         int               `json:"step_number"`
	Title              string            `json:"title"`
	Description        string            `json:"description"`
}
