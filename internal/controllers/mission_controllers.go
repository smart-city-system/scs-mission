package http

import (
	"io"
	"net/http"
	"scs-guard/internal/dto"
	services "scs-guard/internal/services"
	"scs-guard/pkg/validation"

	"github.com/labstack/echo/v4"
)

// MissionHandler handles mission-related HTTP requests
// @Description Mission handler for managing incident guidance and assignments
type MissionHandler struct {
	svc services.MissionService
}

// NewMissionHandler constructor
func NewMissionHandler(svc services.MissionService) *MissionHandler {
	return &MissionHandler{svc: svc}
}

// GetAssignments retrieves mission assignments for a user
// @Summary Get user mission assignments
// @Description Retrieve all mission assignments for the authenticated user
// @Tags missions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Success 200 {object} middleware.SuccessResponse{data=[]models.IncidentGuidance} "List of mission assignments"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /api/v1/missions/me [get]
func (h *MissionHandler) GetAssignments(userID string) echo.HandlerFunc {
	return func(c echo.Context) error {
		assignments, err := h.svc.GetAssignments(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, assignments)
	}
}

// CompleteStep marks a mission step as completed
// @Summary Complete a mission step
// @Description Mark a specific step in a mission guidance as completed
// @Tags missions
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CompleteMissionDto true "Complete mission request"
// @Success 200 {object} middleware.SuccessResponse{data=string} "Step completed successfully"
// @Failure 400 {object} errors.ErrorResponse "Bad request - validation error"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 404 {object} errors.ErrorResponse "Mission or step not found"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /api/v1/missions/complete [patch]
func (h *MissionHandler) CompleteStep() echo.HandlerFunc {
	return func(c echo.Context) error {
		var completeMissionDto dto.CompleteMissionDto
		if err := c.Bind(&completeMissionDto); err != nil {
			return err
		}
		if err := validation.ValidateStruct(completeMissionDto); err != nil {
			return err
		}

		err := h.svc.CompleteStep(c.Request().Context(), completeMissionDto)
		if err != nil {
			return err
		}
		return c.JSON(200, "success")
	}
}

// UpdateIncidentInfo uploads media files for an incident
// @Summary Upload incident media files
// @Description Upload image or video files to document an incident
// @Tags missions
// @Accept multipart/form-data
// @Produce json
// @Security BearerAuth
// @Param incident_id formData string true "Incident ID"
// @Param files formData file true "Media files (images or videos, max 10MB each)"
// @Success 200 {object} middleware.SuccessResponse{data=string} "Files uploaded successfully"
// @Failure 400 {object} errors.ErrorResponse "Bad request - invalid file type or size"
// @Failure 401 {object} errors.ErrorResponse "Unauthorized"
// @Failure 500 {object} errors.ErrorResponse "Internal server error"
// @Router /api/v1/missions/update [put]
func (h *MissionHandler) UpdateIncidentInfo() echo.HandlerFunc {
	return func(c echo.Context) error {
		form, err := c.MultipartForm()
		if err != nil {
			return err
		}
		fileHeaders := form.File["files"]
		incidentID := form.Value["incident_id"]

		const maxSize = 10 * 1024 * 1024 // 10MB
		allowedTypes := map[string]bool{
			"image/": true,
			"video/": true,
		}
		var validFiles []map[string]interface{}
		for _, fileHeader := range fileHeaders {
			if fileHeader.Size > maxSize {
				return echo.NewHTTPError(400, "file size exceeds 10MB")
			}
			file, err := fileHeader.Open()
			if err != nil {
				return echo.NewHTTPError(400, "cannot open file")
			}

			defer file.Close()
			buf := make([]byte, 512)
			n, _ := file.Read(buf)
			mimeType := http.DetectContentType(buf[:n])
			if seeker, ok := file.(io.Seeker); ok {
				seeker.Seek(0, io.SeekStart)
			}
			valid := false
			for prefix := range allowedTypes {
				if len(mimeType) >= len(prefix) && mimeType[:len(prefix)] == prefix {
					valid = true
					break
				}
			}
			if !valid {
				return echo.NewHTTPError(400, "invalid file type: only image and video allowed")
			}
			validFiles = append(validFiles, map[string]interface{}{
				"file":      file,
				"mime_type": mimeType,
				"file_name": incidentID[0] + "/" + fileHeader.Filename,
				"file_size": fileHeader.Size,
			})

		}
		err = h.svc.UpdateIncidentInfo(c.Request().Context(), incidentID[0], validFiles)
		if err != nil {
			return err
		}

		return c.JSON(200, "success")
	}
}
