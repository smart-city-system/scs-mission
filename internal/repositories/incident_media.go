package repositories

import (
	"context"
	"fmt"
	"scs-guard/internal/models"

	"gorm.io/gorm"
)

// IncidentMediaRepository is a repository for incident media
type IncidentMediaRepository struct {
	db *gorm.DB
}

func NewIncidentMediaRepository(db *gorm.DB) *IncidentMediaRepository {
	return &IncidentMediaRepository{db: db}
}

// Create creates a new incident media
func (r *IncidentMediaRepository) BatchCreate(ctx context.Context, incidentMedias []models.IncidentMedia) error {
	if err := r.db.WithContext(ctx).Create(incidentMedias).Error; err != nil {
		return fmt.Errorf("failed to create incident medias: %w", err)
	}
	return nil
}
