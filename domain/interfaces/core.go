package interfaces

import "time"

// TimeProvider provides current time for dependency injection and testing
type TimeProvider interface {
	Now() time.Time
}

// IDGenerator generates unique identifiers
type IDGenerator interface {
	Generate() string
}

// Validator provides validation capabilities
type Validator interface {
	Validate(interface{}) error
}

// Logger provides logging capabilities (interface only, implementation in shared-logging)
type Logger interface {
	Debug(message string, fields map[string]interface{})
	Info(message string, fields map[string]interface{})
	Warning(message string, fields map[string]interface{})
	Error(message string, fields map[string]interface{})
	Critical(message string, fields map[string]interface{})
}
