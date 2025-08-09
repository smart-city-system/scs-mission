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
	UserRepo                 *repositories.UserRepository
	// Services
	MissionService *services.MissionService
}

// NewContainer creates a new dependency container with all repositories and services
func NewContainer(db *gorm.DB, minioClient *minio_client.MinioClient) *Container {
	// Initialize repositories
	incidentGuidanceRepo := repositories.NewIncidentGuidanceRepository(db)
	incidentGuidanceStepRepo := repositories.NewIncidentGuidanceStepRepository(db)
	userRepo := repositories.NewUserRepository(db)
	// Initialize services

	missionService := services.NewMissionService(*incidentGuidanceRepo, *incidentGuidanceStepRepo, *minioClient)

	return &Container{
		// Repositories
		IncidentGuidanceRepo:     incidentGuidanceRepo,
		IncidentGuidanceStepRepo: incidentGuidanceStepRepo,
		UserRepo:                 userRepo,
		// Services
		MissionService: missionService,
	}
}
