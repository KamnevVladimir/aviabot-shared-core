package customTypes

import (
	"fmt"
)

// BusinessRuleError представляет ошибку нарушения бизнес-правила
type BusinessRuleError struct {
	rule    string
	details string
}

// NewBusinessRuleError создает новую ошибку бизнес-правила
func NewBusinessRuleError(rule, details string) BusinessRuleError {
	return BusinessRuleError{
		rule:    rule,
		details: details,
	}
}

// Error возвращает строковое представление ошибки бизнес-правила
func (e BusinessRuleError) Error() string {
	if e.rule == "" {
		if e.details != "" {
			return fmt.Sprintf("business rule violation: %s", e.details)
		}
		return "business rule violation"
	}

	if e.details != "" {
		return fmt.Sprintf("business rule violation '%s': %s", e.rule, e.details)
	}

	return fmt.Sprintf("business rule violation '%s'", e.rule)
}

// Type возвращает тип ошибки
func (e BusinessRuleError) Type() string {
	return "BusinessRuleError"
}

// Rule возвращает название нарушенного бизнес-правила
func (e BusinessRuleError) Rule() string {
	return e.rule
}

// Details возвращает детали нарушения бизнес-правила
func (e BusinessRuleError) Details() string {
	return e.details
}
