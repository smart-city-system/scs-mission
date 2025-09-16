package repositories

import (
	"context"
	"fmt"
	"scs-guard/internal/models"

	"gorm.io/gorm"
)

type IncidentRepository struct {
	db *gorm.DB
}

func NewIncidentRepository(db *gorm.DB) *IncidentRepository {
	return &IncidentRepository{db: db}
}

func (r *IncidentRepository) CreateIncident(ctx context.Context, Incident *models.Incident) (*models.Incident, error) {
	if err := r.db.WithContext(ctx).Create(Incident).Error; err != nil {
		return nil, fmt.Errorf("failed to create Incident: %w", err)
	}
	return Incident, nil
}
func (r *IncidentRepository) GetIncidents(ctx context.Context) ([]models.Incident, error) {
	var Incidents []models.Incident
	if err := r.db.WithContext(ctx).Find(&Incidents).Error; err != nil {
		return nil, fmt.Errorf("failed to get Incidents: %w", err)
	}
	return Incidents, nil
}

func (r *IncidentRepository) GetIncidentByID(ctx context.Context, id string) (*models.Incident, error) {
	var Incident models.Incident
	if err := r.db.WithContext(ctx).Preload("IncidentGuidance").
		Preload("IncidentGuidance.IncidentGuidanceSteps").
		Preload("IncidentGuidance.Assignee").
		Preload("IncidentGuidance.Assigner").First(&Incident, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get Incident: %w", err)
	}
	return &Incident, nil
}
