# Gift Buyer - автоматическая покупка Telegram подарков

[![Language: Russian](https://img.shields.io/badge/Language-Русский-blue)](#русский) [![Language: English](https://img.shields.io/badge/Language-English-green)](#english) [![Telegram](https://img.shields.io/badge/Telegram-@chiefssq-blue?logo=telegram)](https://t.me/cheifssq)

## 📑 Содержание

### Русский
- [📖 Описание](#-описание)
- [⚡ Преимущества](#-преимущества)
- [🚀 Быстрый старт](#-быстрый-старт)
  - [Установка зависимостей](#установка-зависимостей)
  - [Настройка переменных окружения](#настройка-переменных-окружения)
  - [Сборка и запуск](#сборка-и-запуск)
- [🛠 Разработка](#-разработка)
  - [Доступные команды](#доступные-команды)
  - [Локальная проверка перед коммитом](#локальная-проверка-перед-коммитом)
- [📁 Структура проекта](#-структура-проекта)
- [🔒 Безопасность](#-безопасность)
- [📊 Тестирование](#-тестирование)
- [🚀 Деплой](#-деплой)
- [📝 Логирование](#-логирование)
- [🤝 Вклад в проект](#-вклад-в-проект)
- [📄 Лицензия](#-лицензия)
- [⚙️ Подробная конфигурация](#️-подробная-конфигурация)
  - [🔧 Telegram настройки](#-telegram-настройки)
  - [🎯 Критерии покупки](#-критерии-покупки)
  - [👤 Получатели подарков](#-получатели-подарков)
  - [⏱️ Частота проверки](#️-частота-проверки)
  - [🚀 Параметры производительности](#-параметры-производительности)
  - [🧪 Тестовый режим](#-тестовый-режим)
  - [🔒 Глобальные ограничители](#-глобальные-ограничители)
- [📋 Полный пример конфигурации](#-полный-пример-конфигурации)

### English
- [📖 Description](#-description)
- [⚡ Advantages](#-advantages)
- [🚀 Quick Start](#-quick-start-1)
- [🛠 Development](#-development)
- [📁 Project Structure](#-project-structure)
- [🔒 Security](#-security)
- [📊 Testing](#-testing)
- [🚀 Deploy](#-deploy)
- [📝 Logging](#-logging)
- [🤝 Contributing](#-contributing)
- [📄 License](#-license)
- [⚙️ Detailed Configuration](#️-detailed-configuration)

---

## Русский

### 📖 Описание

**Gift Buyer** — софт для автоматической покупки Star Gifts в Telegram. Программа непрерывно мониторит доступные подарки, проверяет их соответствие заданным критериям и автоматически покупает подходящие варианты с возможностью приоритизации.

### ⚡ Преимущества

- **🎯 Приоритизация подарков** — настраиваемые критерии позволяют отправлять разные подарки разным получателям
- **🔗 Поддержка тегов** — возможность использовать теги пользователей и каналов вместо ID
- **🌐 Универсальный доступ** — не требуется быть в контактах с пользователем или подписываться на канал
- **⚡ Высокая скорость** — настраиваемый тикер для мониторинга и мгновенной реакции
- **🎛️ Точная фильтрация** — настраиваемые критерии по цене, количеству и лимитам
- **⚙️ Параллельная обработка** — одновременная покупка нескольких подарков
- **🚀 Оптимизированная производительность** — рекомендуемые настройки для максимальной эффективности
- **📱 Уведомления** — мгновенные оповещения в Telegram при необходимости
- **💾 Кэширование** — сохранение состояния между перезапусками
- **🛡️ Безопасность** — graceful shutdown и обработка ошибок
- **🔄 Автоматический реконнект** — умное переподключение при критических ошибках API
- **📱 Логирование в Telegram** — все ошибки и статусы покупок отправляются в Telegram бот
- **⏱️ Контролируемые таймауты** — 10-минутный таймаут для ввода кода с автоматическим отключением
- **🛡️ Устойчивость к сбоям** — пауза мониторинга во время переподключения с последующим восстановлением

### 🚀 Быстрый старт

### Установка зависимостей
```bash
go mod download
```

### Настройка переменных окружения
### Сборка и запуск
```bash
go build -o gift-buyer cmd/main.go
./gift-buyer
```

## 🛠 Разработка

### Доступные команды
```bash
go mod tidy              # Обновить зависимости
go test ./...            # Запустить тесты
go test -v ./...         # Запустить тесты с подробным выводом
go test -cover ./...     # Запустить тесты с покрытием
go vet ./...             # Проверить код статическим анализатором
go fmt ./...             # Отформатировать код
go build -o gift-buyer cmd/main.go  # Собрать бинарный файл
go run cmd/main.go       # Запустить без сборки
```

### Локальная проверка перед коммитом
```bash
go test ./... && go vet ./... && go fmt ./...
```

## 📁 Структура проекта

```
├── cmd/                    # Точки входа приложения
├── internal/              # Внутренняя логика
│   ├── config/           # Конфигурация
│   └── service/          # Бизнес-логика
├── pkg/                  # Переиспользуемые пакеты
├── .github/workflows/    # GitHub Actions
├── docs/                 # Документация
├── go.mod               # Зависимости Go
├── .golangci.yml        # Конфигурация линтера
└── env.example          # Пример переменных окружения
```

## 🔒 Безопасность

Проект использует:
- Криптографически стойкие генераторы случайных чисел
- Безопасные права доступа к файлам (0600)
- Проверку безопасности с помощью gosec
- Обработку всех ошибок

## 📊 Тестирование

```bash
go test ./...           # Запустить все тесты
go test -cover ./...    # Тесты с отчетом о покрытии
```

## 🚀 Деплой

### Ручная сборка
```bash
# Сборка для всех платформ
GOOS=linux GOARCH=amd64 go build -o gift-buyer-linux cmd/main.go
GOOS=windows GOARCH=amd64 go build -o gift-buyer-windows.exe cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o gift-buyer-macos cmd/main.go
```

## 📝 Логирование

Уровни логирования настраиваются через переменную `LOG_LEVEL`:
- `debug` - Подробная отладочная информация
- `info` - Общая информация (по умолчанию)
- `warn` - Предупреждения
- `error` - Только ошибки

### 📱 Уведомления в Telegram

При настроенном боте все важные события автоматически отправляются в Telegram:

**Уведомления о покупках:**
- ✅ Успешная покупка всех подарков
- ⚠️ Частичная покупка (часть подарков куплена)
- ❌ Неудачная покупка с деталями ошибок

**Системные уведомления:**
- 🔄 Начало процесса переподключения
- ✅ Успешное переподключение
- ❌ Критические ошибки API
- ⏱️ Таймауты аутентификации

### 🔄 Система автоматического переподключения

Программа автоматически обнаруживает и обрабатывает критические ошибки API:

**Критические ошибки, требующие переподключения:**
- `AUTH_KEY_UNREGISTERED` - ключ аутентификации не зарегистрирован
- `CONNECTION_NOT_INITED` - соединение не инициализировано
- `SESSION_REVOKED` - сессия отозвана

**Процесс переподключения:**
1. 🔍 Обнаружение критической ошибки API
2. ⏸️ Пауза мониторинга подарков
3. 🔄 Переподключение к Telegram API
4. ⏱️ Таймаут 10 минут для ввода кода (при необходимости)
5. ▶️ Возобновление мониторинга подарков
6. 🚫 Автоматическое отключение при неудаче

**Важно:** Если переподключение не удается в течение 10 минут, программа автоматически завершается для предотвращения бесконечных попыток.

## 🤝 Вклад в проект

1. Форкните репозиторий
2. Создайте ветку для фичи (`git checkout -b feature/amazing-feature`)
3. Запустите `go test ./... && go vet ./...` для проверки
4. Закоммитьте изменения (`git commit -m 'Add amazing feature'`)
5. Запушьте ветку (`git push origin feature/amazing-feature`)
6. Создайте Pull Request

## 📄 Лицензия

Этот проект распространяется под лицензией MIT.

### 🚀 Быстрый запуск

1. **Скачайте и соберите проект:**
   ```bash
   git clone <repository-url>
   cd gift-buyer
   go build -o gift-buyer cmd/main.go
   ```

2. **Настройте конфигурацию:**
   ```bash
   cp internal/config/config_example.json internal/config/config.json
   # Отредактируйте config.json с вашими данными
   ```

3. **Запустите программу:**
   ```bash
   ./gift-buyer
   ```

### ⚠️ Важные новые возможности

**🎯 Приоритизация подарков:**
- Возможность настройки разных критериев для разных получателей
- Каждый критерий может иметь свой список типов получателей (ReceiverType)
- Полная рандомизация заменена на контролируемое распределение

**🔗 Поддержка тегов:**
- Можно использовать теги пользователей (@username) и каналов (@channelname)
- Не требуется знать точные ID пользователей и каналов
- Система автоматически разрешает теги в ID

**🌐 Универсальный доступ:**
- Не требуется быть в контактах с пользователем для отправки подарка
- Не требуется быть подписанным на канал или быть его администратором
- Полная свобода в выборе получателей

**🔄 Автоматическое переподключение:**
- Программа автоматически переподключается при критических ошибках API
- Мониторинг приостанавливается во время переподключения
- При неудаче переподключения программа автоматически завершается

**⏱️ Контролируемые таймауты:**
- 10-минутный таймаут для ввода кода аутентификации
- Автоматическое завершение при превышении таймаута

**📱 Уведомления в Telegram:**
- Настройте бота для получения уведомлений об ошибках
- Мгновенные уведомления о статусе покупок
- Информация о процессе переподключения

### ⚙️ Подробная конфигурация

#### 🔧 Telegram настройки (`tg_settings`)

```json
{
    "tg_settings": {
        "app_id": 12345678,
        "api_hash": "ваш_api_hash",
        "phone": "+1234567890",
        "password": "пароль_2fa",
        "tg_bot_key": "токен_бота",
        "notification_chat_id": 123456789
    }
}
```

- **`app_id`** и **`api_hash`** — обязательные параметры из [my.telegram.org](https://my.telegram.org)
- **`phone`** — номер телефона аккаунта в международном формате
- **`password`** — пароль двухфакторной аутентификации (можно оставить пустым `""` если 2FA отключена)
- **`tg_bot_key`** — токен Telegram бота для уведомлений (**рекомендуется настроить для получения уведомлений об ошибках и статусе покупок**)
- **`notification_chat_id`** — ID чата для отправки уведомлений (ваш user ID)

**💡 Рекомендация:** Настройте Telegram бота для получения важных уведомлений:
- Уведомления о статусе покупок (успех/неудача)
- Критические ошибки API и переподключения
- Таймауты аутентификации и системные сбои

#### 🎯 Критерии покупки (`criterias`)

**🆕 Новая возможность приоритизации!** Теперь каждый критерий может указывать конкретные типы получателей, что позволяет настроить отправку разных подарков разным людям:

```json
{
    "criterias": [
        {
            "min_price": 10,
            "max_price": 100,
            "total_supply": 1000,
            "count": 2,
            "receiver_type": [1]
        },
        {
            "min_price": 500,
            "max_price": 1000,
            "total_supply": 100,
            "count": 1,
            "receiver_type": [0, 2]
        }
    ]
}
```

- **`min_price`** — минимальная цена подарка в звездах
- **`max_price`** — максимальная цена подарка в звездах  
- **`total_supply`** — максимальный общий тираж подарка
- **`count`** — количество подарков для покупки по этому критерию
- **`receiver_type`** — массив типов получателей для этого критерия:
  - `0` — отправить себе
  - `1` — отправить другому пользователю
  - `2` — отправить в канал/супергруппу

**Пример приоритизации:**
- Дешевые подарки (10-100 звезд) отправляются только пользователям (`receiver_type: [1]`)
- Дорогие подарки (500-1000 звезд) отправляются себе или в каналы (`receiver_type: [0, 2]`)

#### 👤 Получатели подарков (`receiver`)

**🆕 Поддержка тегов!** Теперь можно использовать теги пользователей и каналов вместо ID:

```json
{
    "receiver": {
        "user_receiver_id": ["@username1", "@username2", "username3"],
        "channel_receiver_id": ["@channel1", "@channel2", "channel3"]
    }
}
```

**Форматы идентификаторов:**
- **Пользователи:** `@username` (тег)
- **Каналы:** `@channelname` (тег)

**Как это работает:**
- При покупке система использует `receiver_type` из критерия для определения типа получателя
- В зависимости от типа выбирается случайный тег из соответствующего массива
- Система автоматически разрешает теги в ID при необходимости
- **Не требуется быть в контактах с пользователем или подписываться на канал**

**Преимущества новой системы:**
- Полный контроль над тем, кому какие подарки отправляются
- Возможность использования удобных тегов вместо сложных ID
- Универсальный доступ без необходимости добавления в контакты

#### ⏱️ Частота проверки (`ticker`)

```json
{
    "ticker": 2.0
}
```

Интервал между проверками новых подарков в секундах (по умолчанию 2.0 секунды)

#### 🚀 Параметры производительности

**⚡ Рекомендуемые настройки для оптимальной производительности:**

```json
{
    "retry_count": 5,
    "retry_delay": 2.5,
    "concurrency_gift_count": 10,
    "concurrent_operations": 300,
    "rpc_rate_limit": 30
}
```

**📋 Рекомендации по настройкам:**
- **`retry_count`**: **5+** попыток для максимальной надежности
- **`retry_delay`**: **2-3 секунды** для стабильности API
- **`rpc_rate_limit`**: **максимум 30 RPS** для соблюдения лимитов Telegram

**Описание параметров:**
- **`retry_count`** — количество попыток повтора при неудачной покупке (рекомендуется 5+)
- **`retry_delay`** — задержка между попытками повтора в секундах (рекомендуется 2-3 секунды)
- **`concurrency_gift_count`** — максимальное количество подарков, обрабатываемых одновременно
- **`concurrent_operations`** — максимальное количество одновременных операций
- **`rpc_rate_limit`** — лимит RPC запросов в секунду для Telegram API (рекомендуется не более 30 RPS)

**Оптимизация производительности:**
- Система использует асинхронную обработку повторов с настраиваемой задержкой
- RPC запросы ограничены для соблюдения лимитов Telegram
- Параллельная обработка операций для максимальной скорости
- Интеллектуальная система повторов с прогрессивной задержкой

#### 🧪 Тестовый режим (`test_mode`)

```json
{
    "test_mode": true
}
```

В тестовом режиме:
- **Не учитывается** общий тираж (`total_supply`)
- **Не учитывается** капитализация
- **Поле лимитированности должно быть отрицательным** для покупки нелимитированных подарков
- Подарки покупаются без реальных ограничений

#### 🔒 Глобальные ограничители

```json
{
    "total_star_cap": 10000,
    "max_buy_count": 5,
    "limited_status": false
}
```

- **`total_star_cap`** — максимальная капитализация подарка в звездах
- **`max_buy_count`** — глобальный лимит покупок
- **`limited_status`** — фильтр по статусу лимитированности подарков (true = только лимитированные, false = все подарки)

**Пример работы глобального ограничителя:**
Если у вас есть критерии на покупку 3+2+1=6 подарков, но `max_buy_count: 4`, то купится только 4 подарка (в порядке появления).

### 📋 Полный пример конфигурации

```json
{
    "logger_level": "info",
    "soft_config": {
        "tg_settings": {
            "app_id": 12345678,
            "api_hash": "ваш_api_hash",
            "phone": "+1234567890",
            "password": "",
            "tg_bot_key": "",
            "notification_chat_id": 123456789
        },
        "criterias": [
            {
                "min_price": 10,
                "max_price": 50,
                "total_supply": 1000,
                "count": 3,
                "receiver_type": [1]
            },
            {
                "min_price": 100,
                "max_price": 500,
                "total_supply": 500,
                "count": 2,
                "receiver_type": [0, 2]
            }
        ],
        "total_star_cap": 5000,
        "receiver": {
            "user_receiver_id": ["@username1", "@username2", "123456789"],
            "channel_receiver_id": ["@channel1", "@channel2"]
        },
        "test_mode": false,
        "max_buy_count": 4,
        "ticker": 2.0,
        "retry_count": 5,
        "retry_delay": 2.5,
        "limited_status": false,
        "concurrency_gift_count": 10,
        "concurrent_operations": 300,
        "rpc_rate_limit": 30
    }
}
```

### 📋 Требования

- Go 1.23.4+
- Telegram аккаунт с API ключами(https://my.telegram.org/apps)
- Telegram бот для уведомлений (**рекомендуется**)

### 🤖 Настройка Telegram бота (рекомендуется)

Для получения уведомлений об ошибках и статусе покупок настройте Telegram бота:

1. **Создайте бота:**
   - Напишите [@BotFather](https://t.me/BotFather)
   - Используйте команду `/newbot`
   - Следуйте инструкциям для создания бота
   - Скопируйте полученный токен

2. **Получите ваш User ID:**
   - Напишите [@userinfobot](https://t.me/userinfobot)
   - Скопируйте ваш User ID

3. **Настройте конфигурацию:**
   ```json
   {
     "tg_bot_key": "1234567890:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
     "notification_chat_id": 123456789
   }
   ```

4. **Запустите чат с ботом:**
   - Найдите вашего бота в Telegram
   - Нажмите "Запустить" или отправьте `/start`

После настройки вы будете получать уведомления о:
- ✅/❌ Статусе покупок подарков
- 🔄 Процессе переподключения
- ⚠️ Критических ошибках API
- ⏱️ Таймаутах аутентификации

---

## English

### 📖 Description

**Gift Buyer** is a modern automated system for purchasing Star Gifts in Telegram. The program continuously monitors available gifts, validates them against configured criteria, and automatically purchases eligible options with prioritization capabilities.

### ⚡ Advantages

- **🎯 Gift Prioritization** — configurable criteria allow sending different gifts to different recipients
- **🔗 Tag Support** — ability to use user and channel tags instead of IDs
- **🌐 Universal Access** — no need to be in contacts with users or subscribe to channels
- **🚀 High Speed** — instant reaction to new gift appearances
- **🎛️ Precise Filtering** — configurable criteria for price, quantity, and limits
- **⚙️ Parallel Processing** — simultaneous purchase of multiple gifts
- **🚀 Optimized Performance** — recommended settings for maximum efficiency
- **📱 Notifications** — instant Telegram alerts
- **💾 Caching** — state persistence between restarts
- **🛡️ Security** — graceful shutdown and error handling
- **🔄 Automatic Reconnect** — smart reconnection on critical API errors
- **📱 Telegram Logging** — all errors and purchase statuses are sent to Telegram bot
- **⏱️ Controlled Timeouts** — 10-minute timeout for entering code with automatic deactivation
- **🛡️ Fault Tolerance** — pause monitoring during reconnection with subsequent restoration

### 🚀 Quick Start

### Install dependencies
```bash
go mod download
```

### Set up environment variables
Copy the example environment file:
```bash
cp env.example .env
```

Fill in the `.env` file with your data:
- `TG_APP_ID` - Telegram application ID
- `TG_API_HASH` - API Hash from Telegram
- `TG_PHONE` - Phone number
- `TG_PASSWORD` - Password (optional)

### Build and run
```bash
go build -o gift-buyer cmd/main.go
./gift-buyer
```

## 🛠 Development

### Available commands
```bash
go mod tidy              # Update dependencies
go test ./...            # Run tests
go test -v ./...         # Run tests with verbose output
go test -cover ./...     # Run tests with coverage
go vet ./...             # Check code with static analyzer
go fmt ./...             # Format code
go build -o gift-buyer cmd/main.go  # Build binary file
go run cmd/main.go       # Run without building
```

### Local check before commit
```bash
go test ./... && go vet ./... && go fmt ./...
```

## 📁 Project Structure

```
├── cmd/                    # Application entry points
├── internal/              # Internal logic
│   ├── config/           # Configuration
│   └── service/          # Business logic
├── pkg/                  # Reusable packages
├── .github/workflows/    # GitHub Actions
├── docs/                 # Documentation
├── go.mod               # Go dependencies
├── .golangci.yml        # Linter configuration
└── env.example          # Environment variables example
```

## 🔒 Security

The project uses:
- Cryptographically secure random number generators
- Secure file permissions (0600)
- Security checks with gosec
- Handling of all errors

## 📊 Testing

```bash
go test ./...           # Run all tests
go test -cover ./...    # Tests with coverage report
```

Coverage report is saved to `coverage.html`.

## 🚀 Deploy

### Automatic deploy
When pushing to `main` branch automatically:
1. All tests are run
2. Binary files are built for all platforms
3. Release is created with artifacts

### Manual build
```bash
# Build for all platforms
GOOS=linux GOARCH=amd64 go build -o gift-buyer-linux cmd/main.go
GOOS=windows GOARCH=amd64 go build -o gift-buyer-windows.exe cmd/main.go
GOOS=darwin GOARCH=amd64 go build -o gift-buyer-macos cmd/main.go
```

## 📝 Logging

Logging levels are configured via `LOG_LEVEL` variable:
- `debug` - Detailed debug information
- `info` - General information (default)
- `warn` - Warnings
- `error` - Errors only

### 📱 Telegram Notifications

When a bot is configured, all important events are automatically sent to Telegram:

**Purchase Notifications:**
- ✅ Successful purchase of all gifts
- ⚠️ Partial purchase (some gifts bought)
- ❌ Failed purchase with error details

**System Notifications:**
- 🔄 Reconnection process started
- ✅ Successful reconnection
- ❌ Critical API errors
- ⏱️ Authentication timeouts

### 🔄 Automatic Reconnection System

The program automatically detects and handles critical API errors:

**Critical errors requiring reconnection:**
- `AUTH_KEY_UNREGISTERED` - authentication key not registered
- `CONNECTION_NOT_INITED` - connection not initialized  
- `SESSION_REVOKED` - session revoked

**Reconnection process:**
1. 🔍 Critical API error detection
2. ⏸️ Pause gift monitoring
3. 🔄 Reconnect to Telegram API
4. ⏱️ 10-minute timeout for code entry (if needed)
5. ▶️ Resume gift monitoring
6. 🚫 Automatic shutdown on failure

**Important:** If reconnection fails within 10 minutes, the program automatically terminates to prevent infinite retry loops.

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Run `go test ./... && go vet ./...` for checks
4. Commit your changes (`git commit -m 'Add amazing feature'`)
5. Push the branch (`git push origin feature/amazing-feature`)
6. Create a Pull Request

## 📄 License

This project is distributed under the MIT License.

### ⚙️ Detailed Configuration

#### 🔧 Telegram Settings

**⚠️ IMPORTANT:** Telegram settings are loaded from environment variables, NOT from JSON file!

Required environment variables:
```bash
export TG_APP_ID=12345678
export TG_API_HASH=your_api_hash
export TG_PHONE=+1234567890
```

Optional environment variables:
```bash
export TG_PASSWORD=2fa_password           # Only if 2FA is enabled
export TG_BOT_KEY=bot_token              # For notifications
export TG_NOTIFICATION_CHAT_ID=123456789 # Chat ID for notifications
export DEVICE_MODEL="MacBook Pro M1 Pro" # Device model
export SYSTEM_VERSION="macOS 15.3.1"     # System version
export APP_VERSION="11.9 (272031) APP_STORE" # App version
export SYSTEM_LANG_CODE=en               # System language
export LANG_CODE=en                      # Interface language
export LANG_PACK=en                      # Language pack
```

In JSON file, the `tg_settings` section is ignored - all values are taken from environment variables.

**Parameter descriptions:**
- **`TG_APP_ID`** and **`TG_API_HASH`** — required parameters from [my.telegram.org](https://my.telegram.org)
- **`TG_PHONE`** — account phone number in international format
- **`TG_PASSWORD`** — two-factor authentication password (can be omitted if 2FA is disabled)
- **`TG_BOT_KEY`** — Telegram bot token for notifications (**recommended for error notifications and purchase status**)
- **`TG_NOTIFICATION_CHAT_ID`** — chat ID for sending notifications (your user ID)

**💡 Recommendation:** Configure a Telegram bot for important notifications:
- Purchase status notifications (success/failure)
- Critical API errors and reconnections
- Authentication timeouts and system failures

#### 🎯 Purchase Criteria (`criterias`)

**🆕 New prioritization feature!** Each criteria can now specify specific recipient types, allowing you to configure sending different gifts to different people:

```json
{
    "criterias": [
        {
            "min_price": 10,
            "max_price": 100,
            "total_supply": 1000,
            "count": 2,
            "receiver_type": [1]
        },
        {
            "min_price": 500,
            "max_price": 1000,
            "total_supply": 100,
            "count": 1,
            "receiver_type": [0, 2]
        }
    ]
}
```

- **`min_price`** — minimum gift price in stars
- **`max_price`** — maximum gift price in stars
- **`total_supply`** — maximum total gift supply
- **`count`** — number of gifts to purchase for this criteria
- **`receiver_type`** — array of recipient types for this criteria:
  - `0` — send to yourself
  - `1` — send to another user
  - `2` — send to channel/supergroup

**Prioritization example:**
- Cheap gifts (10-100 stars) are sent only to users (`receiver_type: [1]`)
- Expensive gifts (500-1000 stars) are sent to yourself or channels (`receiver_type: [0, 2]`)

#### 👤 Gift Recipients (`receiver`)

**🆕 Tag support!** You can now use user and channel tags instead of IDs:

```json
{
    "receiver": {
        "user_receiver_id": ["@username1", "@username2", "123456789"],
        "channel_receiver_id": ["@channel1", "@channel2", "-1001234567890"]
    }
}
```

**Identifier formats:**
- **Users:** `@username` (tag) or `123456789` (ID)
- **Channels:** `@channelname` (tag) or `-1001234567890` (full channel ID)

**How it works:**
- During purchase, the system uses `receiver_type` from criteria to determine recipient type
- Depending on the type, a random ID/tag is selected from the corresponding array
- The system automatically resolves tags to IDs when necessary
- **No need to be in contacts with user or subscribe to channel**

**Advantages of the new system:**
- Full control over which gifts go to which recipients
- Ability to use convenient tags instead of complex IDs
- Universal access without needing to add to contacts

#### ⏱️ Check Frequency (`ticker`)

```json
{
    "ticker": 2.0
}
```

Interval between gift checks in seconds (default 2.0 seconds)

#### 🚀 Performance Parameters

**⚡ Recommended settings for optimal performance:**

```json
{
    "retry_count": 5,
    "retry_delay": 2.5,
    "concurrency_gift_count": 10,
    "concurrent_operations": 300,
    "rpc_rate_limit": 30
}
```

**📋 Setting recommendations:**
- **`retry_count`**: **5+** attempts for maximum reliability
- **`retry_delay`**: **2-3 seconds** for API stability
- **`rpc_rate_limit`**: **maximum 30 RPS** to comply with Telegram limits

**Parameter descriptions:**
- **`retry_count`** — number of retry attempts for failed purchases (recommended 5+)
- **`retry_delay`** — delay between retry attempts in seconds (recommended 2-3 seconds)
- **`concurrency_gift_count`** — maximum number of gifts processed simultaneously
- **`concurrent_operations`** — maximum number of concurrent operations
- **`rpc_rate_limit`** — RPC request rate limit per second for Telegram API (recommended max 30 RPS)

**Performance Optimization:**
- System uses asynchronous retry processing with configurable delay
- RPC requests are limited to comply with Telegram limits
- Parallel processing of operations for maximum speed
- Intelligent retry system with progressive delay

#### 🧪 Test Mode (`test_mode`)

```json
{
    "test_mode": true
}
```

In test mode:
- **Total supply is ignored** (`total_supply`)
- **Capitalization is ignored**
- **Limited field should be negative** to buy unlimited gifts
- Gifts are purchased without real restrictions

#### 🔒 Global Limiters

```json
{
    "total_star_cap": 10000,
    "max_buy_count": 5,
    "limited_status": false
}
```

- **`total_star_cap`** — maximum stars to spend
- **`max_buy_count`** — global purchase limit
- **`limited_status`** — filter by gift limited status (true = only limited gifts, false = all gifts)

**Global limiter example:**
If you have criteria for 3+2+1=6 gifts, but `max_buy_count: 4`, only 4 gifts will be purchased (in order of appearance).

### 📋 Complete Configuration Example

```json
{
    "logger_level": "info",
    "soft_config": {
        "tg_settings": {
            "app_id": 12345678,
            "api_hash": "your_api_hash",
            "phone": "+1234567890",
            "password": "",
            "tg_bot_key": "",
            "notification_chat_id": 123456789
        },
        "criterias": [
            {
                "min_price": 10,
                "max_price": 50,
                "total_supply": 1000,
                "count": 3,
                "receiver_type": [1]
            },
            {
                "min_price": 100,
                "max_price": 500,
                "total_supply": 500,
                "count": 2,
                "receiver_type": [0, 2]
            }
        ],
        "total_star_cap": 5000,
        "receiver": {
            "user_receiver_id": ["@username1", "@username2", "123456789"],
            "channel_receiver_id": ["@channel1", "@channel2"]
        },
        "test_mode": false,
        "max_buy_count": 4,
        "ticker": 2.0,
        "retry_count": 5,
        "retry_delay": 2.5,
        "limited_status": false,
        "concurrency_gift_count": 10,
        "concurrent_operations": 300,
        "rpc_rate_limit": 30
    }
}
```

### 📋 Requirements

- Go 1.23.4+
- Telegram account with API credentials  
- Telegram bot for notifications (**recommended**)

### 🤖 Setting up Telegram Bot (recommended)

To receive error notifications and purchase status updates, configure a Telegram bot:

1. **Create a bot:**
   - Message [@BotFather](https://t.me/BotFather)
   - Use `/newbot` command
   - Follow instructions to create the bot
   - Copy the received token

2. **Get your User ID:**
   - Message [@userinfobot](https://t.me/userinfobot)
   - Copy your User ID

3. **Configure settings:**
   ```json
   {
     "tg_bot_key": "1234567890:AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
     "notification_chat_id": 123456789
   }
   ```

4. **Start chat with bot:**
   - Find your bot in Telegram
   - Click "Start" or send `/start`

After setup, you will receive notifications about:
- ✅/❌ Gift purchase status
- 🔄 Reconnection process
- ⚠️ Critical API errors
- ⏱️ Authentication timeouts

---

## ⚠️ Disclaimer

This software is provided "as is" for educational purposes. Users are responsible for compliance with Telegram's Terms of Service and any financial transactions performed by the software.

## 📄 License

This project is provided as-is for educational and personal use.
