package http

import (
	"io"
	"net/http"
	"scs-guard/internal/dto"
	services "scs-guard/internal/services"
	"scs-guard/pkg/validation"

	"github.com/labstack/echo/v4"
)

// Handler
type MissionHandler struct {
	svc services.MissionService
}

// NewHandler constructor
func NewMissionHandler(svc services.MissionService) *MissionHandler {
	return &MissionHandler{svc: svc}
}

func (h *MissionHandler) GetAssignments(userID string) echo.HandlerFunc {
	return func(c echo.Context) error {
		assignments, err := h.svc.GetAssignments(c.Request().Context(), userID)
		if err != nil {
			return err
		}
		return c.JSON(200, assignments)
	}
}
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
			err = h.svc.UpdateIncidentInfo(c.Request().Context(), incidentID[0], fileHeader.Filename, fileHeader.Size, file)
			if err != nil {
				return err
			}
		}

		return c.JSON(200, "success")
	}
}
