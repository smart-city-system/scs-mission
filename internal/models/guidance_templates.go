package models

type GuidanceTemplate struct {
	Base
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	Category      string         `json:"category"`
	GuidanceSteps []GuidanceStep `json:"guidance_steps" gorm:"foreignKey:GuidanceTemplateID"`
}
