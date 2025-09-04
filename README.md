# Shared Core Library

Общие типы и интерфейсы для всех микросервисов.

## Назначение

- Общие типы ошибок (BusinessRuleError, ValidationError)
- Базовые интерфейсы
- Общие константы и типы
- Переиспользуемые структуры данных

## Содержимое

- `domain/errors/` - типы ошибок
- `domain/interfaces/` - базовые интерфейсы

## Использование

```go
import "github.com/KamnevVladimir/aviabot-shared-core/domain/errors"

// Создание ошибки
err := errors.NewBusinessRuleError("Invalid request")
err := errors.NewValidationError("Missing required field")
```

## Версионирование

- v1.0.1 - текущая версия
- Используется во всех микросервисах
