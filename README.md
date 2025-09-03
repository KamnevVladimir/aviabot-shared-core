# Aviabot Shared Core

Общие типы данных и интерфейсы для микросервисов Aviabot.

## Содержимое

- `domain/errors` - Общие типы ошибок (BusinessRuleError, ValidationError)
- `domain/interfaces` - Базовые интерфейсы (IDGenerator)

## Использование

```go
import (
    "github.com/KamnevVladimir/aviabot-shared-core/domain/errors"
    "github.com/KamnevVladimir/aviabot-shared-core/domain/interfaces"
)
```

## Версионирование

Используется семантическое версионирование (SemVer):
- v1.0.0 - Первая стабильная версия
- v1.1.0 - Новые функции (обратно совместимо)
- v2.0.0 - Breaking changes

## Установка

```bash
go get github.com/KamnevVladimir/aviabot-shared-core@v1.0.0
```
