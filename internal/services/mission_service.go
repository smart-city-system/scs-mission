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
	minioClient              minio_client.MinioClient
}

func NewMissionService(incidentGuidanceRepo repositories.IncidentGuidanceRepository, incidentGuidanceStepRepo repositories.IncidentGuidanceStepRepository, minioClient minio_client.MinioClient) *MissionService {
	// TODO: Pass minioClient as a parameter or initialize here as needed
	return &MissionService{
		incidentGuidanceRepo:     incidentGuidanceRepo,
		incidentGuidanceStepRepo: incidentGuidanceStepRepo,
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

func (s *MissionService) UpdateIncidentInfo(ctx context.Context, incidentID string, fileName string, fileSize int64, file multipart.File) error {
	bucketName := "smart-city" // Change as needed
	objectName := incidentID + "/" + fileName

	_, err := s.minioClient.UploadFile(bucketName, objectName, file, fileSize)
	if err != nil {
		return errors.NewInternalError("failed to upload to minio", err)
	}
	return nil
}
