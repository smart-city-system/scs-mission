package models

import "github.com/google/uuid"

// IncidentMedia represents media files associated with an incident
// @Description Media files (images, videos) attached to incidents for documentation
type IncidentMedia struct {
	Base
	IncidentID uuid.UUID `json:"incident_id" example:"550e8400-e29b-41d4-a716-446655440000" swaggertype:"string" format:"uuid"`
	Incident   *Incident `json:"incident,omitempty" gorm:"foreignKey:IncidentID"`
	MediaType  string    `json:"media_type" gorm:"check:media_type IN ('image', 'video')" example:"image" enums:"image,video"`
	FileUrl    string    `json:"file_url" example:"https://example.com/media/incident_123.jpg"`
	FileSize   int64     `json:"file_size" example:"1024000"`
	FileType   string    `json:"file_type" example:"image/jpeg"`
	FileName   string    `json:"file_name" example:"incident_photo_001.jpg"`
}
