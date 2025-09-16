package models

import "github.com/google/uuid"

type IncidentMedia struct {
	Base
	IncidentID uuid.UUID `json:"incident_id"`
	Incident   *Incident `json:"incident,omitempty" gorm:"foreignKey:IncidentID"`
	MediaType  string    `json:"media_type" gorm:"check:media_type IN ('', 'video')"`
	FileUrl    string    `json:"file_url"`
	FileSize   int64     `json:"file_size"`
	FileType   string    `json:"file_type"`
	FileName   string    `json:"file_name"`
}
