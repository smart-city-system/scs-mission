package repositories

import (
	"context"
	"fmt"
	"scs-guard/internal/models"

	"gorm.io/gorm"
)

type IncidentGuidanceStepRepository struct {
	db *gorm.DB
}

func NewIncidentGuidanceStepRepository(db *gorm.DB) *IncidentGuidanceStepRepository {
	return &IncidentGuidanceStepRepository{db: db}
}
func (r *IncidentGuidanceStepRepository) CreateIncidentGuidanceStep(ctx context.Context, guidance *models.IncidentGuidanceStep) (*models.IncidentGuidanceStep, error) {
	if err := r.db.WithContext(ctx).Create(guidance).Error; err != nil {
		return nil, fmt.Errorf("failed to assign guidance: %w", err)
	}
	return guidance, nil
}

func (r *IncidentGuidanceStepRepository) CreateIncidentGuidanceSteps(ctx context.Context, incidentGuidanceSteps []models.IncidentGuidanceStep) ([]models.IncidentGuidanceStep, error) {
	if err := r.db.WithContext(ctx).Create(incidentGuidanceSteps).Error; err != nil {
		return nil, fmt.Errorf("failed to assign guidance: %w", err)
	}
	return incidentGuidanceSteps, nil
}

func (r *IncidentGuidanceStepRepository) UpdateIncidentGuidanceStep(ctx context.Context, id string, isCompleted bool) error {
	result := r.db.WithContext(ctx).Model(&models.IncidentGuidanceStep{}).Where("id = ?", id).Update("is_completed", isCompleted)
	if result.Error != nil {
		return fmt.Errorf("failed to update guidance step: %w", result.Error)
	}
	return nil
}
func (r *IncidentGuidanceStepRepository) GetIncidentGuidanceStepByID(ctx context.Context, id string) (*models.IncidentGuidanceStep, error) {
	var incidentGuidanceStep models.IncidentGuidanceStep
	if err := r.db.WithContext(ctx).First(&incidentGuidanceStep, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get incident guidance step: %w", err)
	}
	return &incidentGuidanceStep, nil
}
