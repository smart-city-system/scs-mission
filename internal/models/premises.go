package models

import "github.com/google/uuid"

// Premise represents a physical location or building
// @Description Premise entity representing physical locations in the system
type Premise struct {
	Base
	Name            string     `json:"name" example:"Main Building"`
	Address         string     `json:"address" example:"123 Main Street, City, Country"`
	ParentPremiseID *uuid.UUID `json:"parent_premise_id,omitempty" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	ParentPremise   *Premise   `json:"parent_premise,omitempty" gorm:"foreignKey:ParentPremiseID"`
}
