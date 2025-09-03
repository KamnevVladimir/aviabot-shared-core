package customTypes

import (
	"testing"
)

// TestBusinessRuleError_NewBusinessRuleError тестирует создание BusinessRuleError
func TestBusinessRuleError_NewBusinessRuleError(t *testing.T) {
	rule := "test_rule"
	details := "test details"
	
	err := NewBusinessRuleError(rule, details)
	
	if err.Rule() != rule {
		t.Errorf("NewBusinessRuleError() Rule = %v, want %v", err.Rule(), rule)
	}
	
	if err.Details() != details {
		t.Errorf("NewBusinessRuleError() Details = %v, want %v", err.Details(), details)
	}
}

// TestBusinessRuleError_Error тестирует форматирование ошибки бизнес-правила
func TestBusinessRuleError_Error(t *testing.T) {
	tests := []struct {
		name    string
		rule    string
		details string
		want    string
	}{
		{
			name:    "business rule error with rule and details",
			rule:    "user_age_limit",
			details: "user must be at least 18 years old",
			want:    "business rule violation 'user_age_limit': user must be at least 18 years old",
		},
		{
			name:    "business rule error with rule only",
			rule:    "user_age_limit",
			details: "",
			want:    "business rule violation 'user_age_limit'",
		},
		{
			name:    "business rule error with details only",
			rule:    "",
			details: "user must be at least 18 years old",
			want:    "business rule violation: user must be at least 18 years old",
		},
		{
			name:    "business rule error without rule and details",
			rule:    "",
			details: "",
			want:    "business rule violation",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewBusinessRuleError(tt.rule, tt.details)
			got := err.Error()
			if got != tt.want {
				t.Errorf("BusinessRuleError.Error() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestBusinessRuleError_Type тестирует получение типа ошибки
func TestBusinessRuleError_Type(t *testing.T) {
	err := NewBusinessRuleError("test_rule", "test details")
	
	got := err.Type()
	want := "BusinessRuleError"
	
	if got != want {
		t.Errorf("BusinessRuleError.Type() = %v, want %v", got, want)
	}
}

// TestBusinessRuleError_Rule тестирует получение правила ошибки
func TestBusinessRuleError_Rule(t *testing.T) {
	tests := []struct {
		name string
		rule string
	}{
		{
			name: "rule with value",
			rule: "user_age_limit",
		},
		{
			name: "empty rule",
			rule: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewBusinessRuleError(tt.rule, "")
			got := err.Rule()
			if got != tt.rule {
				t.Errorf("BusinessRuleError.Rule() = %v, want %v", got, tt.rule)
			}
		})
	}
}

// TestBusinessRuleError_Details тестирует получение деталей ошибки
func TestBusinessRuleError_Details(t *testing.T) {
	tests := []struct {
		name    string
		details string
	}{
		{
			name:    "details with value",
			details: "user must be at least 18 years old",
		},
		{
			name:    "empty details",
			details: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := NewBusinessRuleError("", tt.details)
			got := err.Details()
			if got != tt.details {
				t.Errorf("BusinessRuleError.Details() = %v, want %v", got, tt.details)
			}
		})
	}
}