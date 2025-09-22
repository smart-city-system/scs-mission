package dto

// CompleteMissionDto represents the request to complete a mission step
// @Description Request payload for completing a mission step
type CompleteMissionDto struct {
	MissionID string `json:"mission_id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440000"`
	StepID    string `json:"step_id" validate:"required" example:"550e8400-e29b-41d4-a716-446655440001"`
}
