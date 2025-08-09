package dto

type CompleteMissionDto struct {
	MissionID string `json:"mission_id" validate:"required"`
	StepID    string `json:"step_id" validate:"required"`
}
