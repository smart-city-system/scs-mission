package container

import (
	repositories "scs-guard/internal/repositories"
	"scs-guard/internal/services"
	minio_client "scs-guard/pkg/minio"

	"gorm.io/gorm"
)

// Container holds all the application dependencies
type Container struct {
	// Repositories
	IncidentGuidanceRepo     *repositories.IncidentGuidanceRepository
	IncidentGuidanceStepRepo *repositories.IncidentGuidanceStepRepository
	IncidentRepo             *repositories.IncidentRepository
	IncidentMediaRepo        *repositories.IncidentMediaRepository
	UserRepo                 *repositories.UserRepository
	// Services
	MissionService *services.MissionService
}

// NewContainer creates a new dependency container with all repositories and services
func NewContainer(db *gorm.DB, minioClient *minio_client.MinioClient) *Container {
	// Initialize repositories
	incidentGuidanceRepo := repositories.NewIncidentGuidanceRepository(db)
	incidentGuidanceStepRepo := repositories.NewIncidentGuidanceStepRepository(db)
	incidentRepo := repositories.NewIncidentRepository(db)
	incidentMediaRepo := repositories.NewIncidentMediaRepository(db)
	userRepo := repositories.NewUserRepository(db)
	// Initialize services

	missionService := services.NewMissionService(*incidentGuidanceRepo, *incidentGuidanceStepRepo, *incidentRepo, *incidentMediaRepo, *minioClient)

	return &Container{
		// Repositories
		IncidentGuidanceRepo:     incidentGuidanceRepo,
		IncidentGuidanceStepRepo: incidentGuidanceStepRepo,
		UserRepo:                 userRepo,
		// Services
		MissionService: missionService,
	}
}
