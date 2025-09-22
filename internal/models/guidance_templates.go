package models

// GuidanceTemplate represents a template for incident guidance procedures
// @Description Template containing step-by-step guidance for handling incidents
type GuidanceTemplate struct {
	Base
	Name          string         `json:"name" example:"Fire Emergency Response"`
	Description   string         `json:"description" example:"Standard procedure for handling fire emergencies"`
	Category      string         `json:"category" example:"Emergency"`
	GuidanceSteps []GuidanceStep `json:"guidance_steps" gorm:"foreignKey:GuidanceTemplateID"`
}
