package customTypes

import (
	"fmt"
)

// ValidationError представляет ошибку валидации с дополнительной информацией о поле
type ValidationError struct {
	field       string
	originalErr error
}

// NewValidationError создает новую ошибку валидации
func NewValidationError(field string, err error) ValidationError {
	return ValidationError{
		field:       field,
		originalErr: err,
	}
}

// Error возвращает строковое представление ошибки валидации
func (e ValidationError) Error() string {
	if e.field == "" {
		if e.originalErr != nil {
			return fmt.Sprintf("validation failed: %s", e.originalErr.Error())
		}
		return "validation failed"
	}

	if e.originalErr != nil {
		return fmt.Sprintf("validation failed for field '%s': %s", e.field, e.originalErr.Error())
	}

	return fmt.Sprintf("validation failed for field '%s'", e.field)
}

// Type возвращает тип ошибки
func (e ValidationError) Type() string {
	return "ValidationError"
}

// Field возвращает поле, которое не прошло валидацию
func (e ValidationError) Field() string {
	return e.field
}

// Unwrap возвращает оригинальную ошибку для использования с errors.Is и errors.As
func (e ValidationError) Unwrap() error {
	return e.originalErr
}
