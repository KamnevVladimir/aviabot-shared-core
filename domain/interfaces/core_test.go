package interfaces

import (
	"testing"
	"time"
)

// TestTimeProvider ensures TimeProvider interface can be implemented
func TestTimeProvider_Interface(t *testing.T) {
	var _ TimeProvider = (*mockTimeProvider)(nil)
}

// TestIDGenerator ensures IDGenerator interface can be implemented  
func TestIDGenerator_Interface(t *testing.T) {
	var _ IDGenerator = (*mockIDGenerator)(nil)
}

// TestValidator ensures Validator interface can be implemented
func TestValidator_Interface(t *testing.T) {
	var _ Validator = (*mockValidator)(nil)
}

// TestLogger ensures Logger interface can be implemented
func TestLogger_Interface(t *testing.T) {
	var _ Logger = (*mockLogger)(nil)
}

// Mock implementations for testing

type mockTimeProvider struct{}

func (m *mockTimeProvider) Now() time.Time {
	return time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
}

type mockIDGenerator struct{}

func (m *mockIDGenerator) Generate() string {
	return "test-id-123"
}

type mockValidator struct{}

func (m *mockValidator) Validate(interface{}) error {
	return nil
}

type mockLogger struct{}

func (m *mockLogger) Debug(message string, fields map[string]interface{})    {}
func (m *mockLogger) Info(message string, fields map[string]interface{})     {}
func (m *mockLogger) Warning(message string, fields map[string]interface{})  {}
func (m *mockLogger) Error(message string, fields map[string]interface{})    {}
func (m *mockLogger) Critical(message string, fields map[string]interface{}) {}

// Integration tests

func TestMockTimeProvider_Now(t *testing.T) {
	provider := &mockTimeProvider{}
	now := provider.Now()
	
	expected := time.Date(2025, 1, 1, 0, 0, 0, 0, time.UTC)
	if !now.Equal(expected) {
		t.Errorf("Now() = %v, want %v", now, expected)
	}
}

func TestMockIDGenerator_Generate(t *testing.T) {
	generator := &mockIDGenerator{}
	id := generator.Generate()
	
	expected := "test-id-123"
	if id != expected {
		t.Errorf("Generate() = %v, want %v", id, expected)
	}
}

func TestMockValidator_Validate(t *testing.T) {
	validator := &mockValidator{}
	err := validator.Validate("test")
	
	if err != nil {
		t.Errorf("Validate() = %v, want nil", err)
	}
}

func TestMockLogger_Methods(t *testing.T) {
	logger := &mockLogger{}
	fields := map[string]interface{}{"key": "value"}
	
	// Test that all methods can be called without panic
	logger.Debug("test", fields)
	logger.Info("test", fields)
	logger.Warning("test", fields)
	logger.Error("test", fields)
	logger.Critical("test", fields)
}