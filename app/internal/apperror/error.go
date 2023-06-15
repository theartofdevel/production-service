package apperror

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"go.opentelemetry.io/otel/trace"
)

const name = "PS"

var (
	ErrInternalSystem = NewAppError(http.StatusInternalServerError, "00100", "internal system error")
	ErrBadRequest     = NewAppError(http.StatusBadRequest, "00101", "bad request")
	ErrValidation     = NewAppError(http.StatusBadRequest, "00102", "validation error")
	ErrNotFound       = NewAppError(http.StatusNotFound, "00103", "not found")
	ErrUnauthorized   = NewAppError(http.StatusUnauthorized, "00104", "unauthorized")
	ErrForbidden      = NewAppError(http.StatusForbidden, "00105", "access forbidden")
)

type ErrorFields map[string]string

type AppError struct {
	Err           error       `json:"-"`
	Message       string      `json:"message,omitempty"`
	Code          string      `json:"code,omitempty"`
	TransportCode int         `json:"-"`
	Fields        ErrorFields `json:"fields,omitempty"`
	TraceID       string      `json:"trace_id,omitempty"`
}

func (e *AppError) WithFields(fields ErrorFields) {
	e.Fields = fields
}

func NewAppError(transportCode int, code, message string) *AppError {
	return &AppError{
		Err:           fmt.Errorf(message),
		Code:          name + "-" + code,
		TransportCode: transportCode,
		Message:       message,
	}
}

func (e *AppError) Error() string {
	err := e.Err.Error()

	if len(e.Fields) > 0 {
		for k, v := range e.Fields {
			err += ", " + k + " " + v
		}
	}

	return err
}

func (e *AppError) Unwrap() error { return e.Err }

func (e *AppError) Marshal(ctx context.Context) []byte {
	if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
		e.TraceID = span.TraceID().String()
	}

	bytes, err := json.Marshal(e)
	if err != nil {
		return nil
	}

	return bytes
}
