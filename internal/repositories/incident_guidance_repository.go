package repositories

import (
	"context"
	"fmt"
	"scs-guard/internal/models"

	"gorm.io/gorm"
)

type IncidentGuidanceRepository struct {
	db *gorm.DB
}

func NewIncidentGuidanceRepository(db *gorm.DB) *IncidentGuidanceRepository {
	return &IncidentGuidanceRepository{db: db}
}
func (r *IncidentGuidanceRepository) CreateIncidentGuidance(ctx context.Context, guidance *models.IncidentGuidance) (*models.IncidentGuidance, error) {
	result := r.db.WithContext(ctx).Create(guidance)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to assign guidance: %w", result.Error)
	}
	return guidance, nil
}
func (r *IncidentGuidanceRepository) GetIncidentGuidanceByIncidentID(ctx context.Context, incidentID string) (*models.IncidentGuidance, error) {
	var incidentGuidance models.IncidentGuidance
	if err := r.db.WithContext(ctx).Preload("Assignee").Preload("Assigner").Preload("Incident").Preload("IncidentGuidanceSteps").First(&incidentGuidance, "incident_id = ?", incidentID).Error; err != nil {
		return nil, fmt.Errorf("failed to get incident guidance: %w", err)
	}
	return &incidentGuidance, nil
}

func (r *IncidentGuidanceRepository) GetIncidentGuidanceByAssigneeID(ctx context.Context, assigneeID string) ([]models.IncidentGuidance, error) {
	var incidentGuidance []models.IncidentGuidance
	if err := r.db.WithContext(ctx).Preload("Assignee").Preload("Assigner").Preload("Incident").Preload("IncidentGuidanceSteps").Find(&incidentGuidance, "assignee_id = ?", assigneeID).Error; err != nil {
		return nil, fmt.Errorf("failed to get incident guidance: %w", err)
	}
	return incidentGuidance, nil
}
