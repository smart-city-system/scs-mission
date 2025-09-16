package services

import (
	"context"
	"mime/multipart"
	"scs-guard/internal/dto"
	"scs-guard/internal/models"
	repositories "scs-guard/internal/repositories"
	"scs-guard/pkg/errors"
	minio_client "scs-guard/pkg/minio"
)

type MissionService struct {
	incidentGuidanceRepo     repositories.IncidentGuidanceRepository
	incidentGuidanceStepRepo repositories.IncidentGuidanceStepRepository
	incidentRepo             repositories.IncidentRepository
	incidentMediaRepo        repositories.IncidentMediaRepository
	minioClient              minio_client.MinioClient
}

func NewMissionService(incidentGuidanceRepo repositories.IncidentGuidanceRepository, incidentGuidanceStepRepo repositories.IncidentGuidanceStepRepository, incidentRepo repositories.IncidentRepository, incidentMediaRepo repositories.IncidentMediaRepository, minioClient minio_client.MinioClient) *MissionService {
	// TODO: Pass minioClient as a parameter or initialize here as needed
	return &MissionService{
		incidentGuidanceRepo:     incidentGuidanceRepo,
		incidentGuidanceStepRepo: incidentGuidanceStepRepo,
		incidentRepo:             incidentRepo,
		incidentMediaRepo:        incidentMediaRepo,
		minioClient:              minioClient,
	}
}

func (s *MissionService) GetAssignments(ctx context.Context, userID string) ([]models.IncidentGuidance, error) {
	assignments, err := s.incidentGuidanceRepo.GetIncidentGuidanceByAssigneeID(ctx, userID)
	if err != nil {
		return nil, errors.NewDatabaseError("get assignments", err)
	}
	return assignments, nil
}

func (s *MissionService) CompleteStep(ctx context.Context, completeMissionDto dto.CompleteMissionDto) error {
	// TODO: Check if the step belongs to the assignment
	stepInfo, err := s.incidentGuidanceStepRepo.GetIncidentGuidanceStepByID(ctx, completeMissionDto.StepID)
	if err != nil {
		return errors.NewDatabaseError("get step", err)
	}
	if stepInfo.IncidentGuidanceID.String() != completeMissionDto.MissionID {
		return errors.NewBadRequestError("step does not belong to the mission")
	}
	if stepInfo.IsCompleted {
		return errors.NewBadRequestError("step already completed")
	}
	s.incidentGuidanceStepRepo.UpdateIncidentGuidanceStep(ctx, completeMissionDto.StepID, true)
	return nil
}

func (s *MissionService) UpdateIncidentInfo(ctx context.Context, incidentID string, validFiles []map[string]interface{}) error {
	incident, err := s.incidentRepo.GetIncidentByID(ctx, incidentID)
	if err != nil {
		return errors.NewBadRequestError("incident not found")
	}
	// Upload files to minio
	var incidentMedias []models.IncidentMedia
	for _, validFile := range validFiles {
		objectName := validFile["file_name"].(string)
		file := validFile["file"].(multipart.File)
		fileSize := validFile["file_size"].(int64)
		fileType := validFile["mime_type"].(string)
		fileInfo, err := s.minioClient.UploadFile(objectName, file, fileSize, fileType)
		if err != nil {
			return err
		}
		incidentMedias = append(incidentMedias, models.IncidentMedia{
			IncidentID: incident.ID,
			FileName:   objectName,
			FileSize:   fileInfo.Size,
			FileUrl:    "http://" + s.minioClient.Endpoint + "/" + s.minioClient.BucketName + "/" + fileInfo.Key,
			MediaType:  getFileType(fileType),
			FileType:   fileType,
		})
	}
	// Create incident media
	if err := s.incidentMediaRepo.BatchCreate(ctx, incidentMedias); err != nil {
		return errors.NewDatabaseError("create incident media", err)
	}

	return nil
}
func getFileType(contentType string) string {
	if contentType == "image/jpeg" || contentType == "image/png" {
		return "image"
	}
	if contentType == "video/mp4" {
		return "video"
	}
	return "other"
}
