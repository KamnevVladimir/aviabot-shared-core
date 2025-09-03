package customTypes

import (
	"errors"
	"testing"
)

// TestValidationError_NewValidationError тестирует создание ValidationError
func TestValidationError_NewValidationError(t *testing.T) {
	originalErr := errors.New("test error")
	field := "test_field"
	
	err := NewValidationError(field, originalErr)
	
	if err.Field() != field {
		t.Errorf("NewValidationError() Field = %v, want %v", err.Field(), field)
	}
	
	if err.Unwrap() != originalErr {
		t.Errorf("NewValidationError() Unwrap = %v, want %v", err.Unwrap(), originalErr)
	}
}

// TestValidationError_Error тестирует форматирование ошибки
func TestValidationError_Error(t *testing.T) {
	tests := []struct {
		name        string
		field       string
		originalErr error
		want        string
	}{
		{
			name:        "validation error with field and original error",
			field:       "user_id",
			originalErr: errors.New("must be positive"),
			want:        "validation failed for field 'user_id': must be positive",
		},
		{
			name:        "validation error with field only",
			field:       "user_id",
			originalErr: nil,
			want:        "validation failed for field 'user_id'",
		},
		{
			name:        "validation error with original error only",
			field:       "",
			originalErr: errors.New("must be positive"),
			want:        "validation failed: must be positive",
		},
		{
			name:        "validation error without field and original error",
			field:       "",
			originalErr: nil,
			want:        "validation failed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewValidationError(tt.field, tt.originalErr)
			got := err.Error()
			if got != tt.want {
				t.Errorf("ValidationError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestValidationError_Type тестирует получение типа ошибки
func TestValidationError_Type(t *testing.T) {
	err := NewValidationError("test_field", errors.New("test error"))
	
	got := err.Type()
	want := "ValidationError"
	
	if got != want {
		t.Errorf("ValidationError.Type() = %v, want %v", got, want)
	}
}

// TestValidationError_Field тестирует получение поля ошибки
func TestValidationError_Field(t *testing.T) {
	tests := []struct {
		name  string
		field string
	}{
		{
			name:  "field with value",
			field: "user_id",
		},
		{
			name:  "empty field",
			field: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewValidationError(tt.field, nil)
			got := err.Field()
			if got != tt.field {
				t.Errorf("ValidationError.Field() = %v, want %v", got, tt.field)
			}
		})
	}
}

// TestValidationError_Unwrap тестирует извлечение оригинальной ошибки
func TestValidationError_Unwrap(t *testing.T) {
	tests := []struct {
		name        string
		originalErr error
	}{
		{
			name:        "with original error",
			originalErr: errors.New("test error"),
		},
		{
			name:        "without original error",
			originalErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewValidationError("test_field", tt.originalErr)
			got := err.Unwrap()
			if got != tt.originalErr {
				t.Errorf("ValidationError.Unwrap() = %v, want %v", got, tt.originalErr)
			}
		})
	}
}