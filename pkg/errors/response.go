package errors

import (
	"time"
)

// ErrorResponse represents the structure of error responses sent to clients
// @Description Standard error response format
type ErrorResponse struct {
	Error     ErrorDetail `json:"error"`
	RequestID string      `json:"request_id,omitempty" example:"req-123456"`
	Timestamp time.Time   `json:"timestamp" example:"2023-01-01T00:00:00Z"`
}

// ErrorDetail contains the error information
// @Description Detailed error information
type ErrorDetail struct {
	Type    ErrorType   `json:"type" example:"VALIDATION_ERROR"`
	Message string      `json:"message" example:"Validation failed"`
	Details interface{} `json:"details,omitempty"`
}

// ValidationError represents validation error details
// @Description Individual field validation error
type ValidationError struct {
	Field   string      `json:"field" example:"email"`
	Message string      `json:"message" example:"Email is required"`
	Value   interface{} `json:"value,omitempty" example:"invalid-email"`
}

// ValidationErrors represents multiple validation errors
// @Description Collection of validation errors
type ValidationErrors []ValidationError

// NewErrorResponse creates a new error response
func NewErrorResponse(appErr *AppError, requestID string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Type:    appErr.Type,
			Message: appErr.Message,
			Details: appErr.Details,
		},
		RequestID: requestID,
		Timestamp: time.Now(),
	}
}

// NewValidationErrorResponse creates a validation error response
func NewValidationErrorResponse(errors ValidationErrors, requestID string) *ErrorResponse {
	return &ErrorResponse{
		Error: ErrorDetail{
			Type:    ErrorTypeValidation,
			Message: "Validation failed",
			Details: errors,
		},
		RequestID: requestID,
		Timestamp: time.Now(),
	}
}
