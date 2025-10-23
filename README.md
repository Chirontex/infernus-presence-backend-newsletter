# Infernus Presence Newsletter Service

Этот сервис предназначен для сайта музыкальной группы Infernus Presence.

Готовый к продакшену Go-бэкенд для подписки на рассылку с использованием базы данных MariaDB.

## Возможности

- ✅ RESTful API для подписки на рассылку
- ✅ Валидация email-адреса
- ✅ Аутентификация клиента по токену
- ✅ MariaDB с миграциями
- ✅ Поддержка Docker (Docker Compose вынесен в отдельный репозиторий)
- ✅ Чистая архитектура (обработчики, сервисы, репозитории)
- ✅ Middleware (логирование, восстановление после ошибок, CORS)
- ✅ Корректное завершение работы
- ✅ Пул соединений
- ✅ Конфигурация через переменные окружения

## Технологический стек

- **Язык**: Go 1.23
- **База данных**: MariaDB 11.5
- **Роутер**: Gorilla Mux
- **Миграции**: golang-migrate
- **Контейнеризация**: Docker

## Структура проекта

```
newsletter-backend/
├── cmd/
│   └── api/
│       └── main.go                 # Точка входа приложения
├── internal/
│   ├── api/
│   │   ├── handlers/              # HTTP-обработчики
│   │   └── middleware/            # Middleware-функции
│   ├── config/                    # Управление конфигурацией
│   ├── database/                  # Подключение к БД и миграции
│   ├── models/                    # Модели данных
│   ├── repository/                # Слой доступа к данным
│   ├── service/                   # Бизнес-логика
│   └── validator/                 # Валидация входных данных
├── migrations/                    # Миграции БД
├── Dockerfile                     # Описание Docker-образа
├── go.mod                         # Зависимости Go
└── Makefile                       # Автоматизация сборки
```

## Документация API

### POST /api/newsletter/subscribe

Подписка на рассылку по email.

**Тело запроса:**
```json
{
  "email": "user@example.com",
  "clientToken": "your-secret-token"
}
```

**Успешный ответ (200):**
```json
{
  "success": true,
  "message": "Успешно подписан на рассылку"
}
```

**Ошибка валидации (400):**
```json
{
  "success": false,
  "error": "неверный формат email"
}
```

**Внутренняя ошибка (500):**
```json
{
  "success": false,
  "error": "Внутренняя ошибка сервера"
}
```

## Быстрый старт с Docker

Для быстрого старта с использованием Docker Compose используйте отдельный репозиторий инфраструктуры, где находится актуальный `docker-compose.yml` и инструкции по запуску всего стека.

### Запуск через Docker (без Compose)

1. Соберите образ:

```bash
docker build -t newsletter-backend:latest .
```

2. Запустите контейнер, передав переменные окружения:

```bash
docker run -d \
  -p 8080:8080 \
  -e CLIENT_TOKEN="your-secret-token" \
  -e DB_HOST="your-db-host" \
  -e DB_USER="your-db-user" \
  -e DB_PASSWORD="your-db-password" \
  -e DB_NAME="newsletter" \
  newsletter-backend:latest
```

> **Примечание:** Контейнер базы данных MariaDB и миграции должны быть запущены отдельно. Используйте внешний репозиторий с docker-compose или настройте БД вручную.

## Локальная разработка (без Docker Compose)

### Требования

- Go 1.23+
- MariaDB 11.5+
- golang-migrate CLI (опционально, для ручных миграций)

### Шаг 1: Установите зависимости

```bash
go mod download
```

### Шаг 2: Настройте базу данных

Создайте базу данных:

```sql
CREATE DATABASE newsletter;
CREATE USER 'newsletter_user'@'localhost' IDENTIFIED BY 'newsletter_password';
GRANT ALL PRIVILEGES ON newsletter.* TO 'newsletter_user'@'localhost';
FLUSH PRIVILEGES;
```

### Шаг 3: Настройте окружение

Скопируйте `.env.example` в `.env` и обновите:

```bash
cp .env.example .env
```

Отредактируйте `.env`:

```env
SERVER_ADDRESS=:8080
CLIENT_TOKEN=your-secret-client-token-here
DB_HOST=localhost
DB_PORT=3306
DB_USER=newsletter_user
DB_PASSWORD=newsletter_password
DB_NAME=newsletter
```

### Шаг 4: Запустите приложение

```bash
go run cmd/api/main.go
```

Приложение автоматически применит миграции при запуске.

### Шаг 5: Сборка для продакшена

```bash
go build -o bin/main cmd/api/main.go
./bin/main
```

## Использование Makefile

```bash
# Сборка приложения
make build

# Запуск приложения
make run

# Запуск тестов
make test

# Очистка артефактов сборки
make clean

# Сборка Docker-образа
make docker-build
```

> **Примечание:** Команды для управления контейнерами через docker-compose удалены. Используйте внешний репозиторий с инфраструктурой для orchestration.

## Схема базы данных

### Таблица emails

| Колонка      | Тип                 | Описание                       |
|------------- |-------------------- |------------------------------- |
| id           | BIGINT UNSIGNED     | Первичный ключ, автоинкремент  |
| email        | VARCHAR(255)        | Email (уникальный)             |
| is_confirmed | TINYINT UNSIGNED    | Статус подтверждения (0 или 1) |
| created_at   | TIMESTAMP           | Дата создания                  |
| updated_at   | TIMESTAMP           | Дата обновления                |
| confirmed_at | TIMESTAMP           | Дата подтверждения             |

**Индексы:**
- `email` — уникальный индекс по email
- `confirmed_email` — составной индекс по (email, is_confirmed)

## Переменные окружения

| Переменная      | Обязательна | По умолчанию | Описание                      |
|---------------- |------------ |------------- |------------------------------ |
| SERVER_ADDRESS  | Нет         | :8080        | Адрес для прослушивания        |
| CLIENT_TOKEN    | Да          | -            | Токен аутентификации клиента   |
| DB_HOST         | Нет         | localhost    | Хост базы данных              |
| DB_PORT         | Нет         | 3306         | Порт базы данных              |
| DB_USER         | Нет         | root         | Пользователь БД               |
| DB_PASSWORD     | Нет         | -            | Пароль БД                     |
| DB_NAME         | Нет         | newsletter   | Имя базы данных               |

## Продакшен-деплой

### Деплой через Docker

1. Соберите образ:
```bash
docker build -t newsletter-backend:latest .
```

2. Запустите с переменными окружения:
```bash
docker run -d \
  -p 8080:8080 \
  -e CLIENT_TOKEN="your-secret-token" \
  -e DB_HOST="your-db-host" \
  -e DB_USER="your-db-user" \
  -e DB_PASSWORD="your-db-password" \
  -e DB_NAME="newsletter" \
  newsletter-backend:latest
```

> **Примечание:** Для orchestration используйте внешний репозиторий с docker-compose или настройте инфраструктуру самостоятельно.

### Деплой через Kubernetes

Создайте ConfigMap для конфигурации:

```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: newsletter-config
data:
  SERVER_ADDRESS: ":8080"
  DB_HOST: "mariadb-service"
  DB_PORT: "3306"
  DB_NAME: "newsletter"
```

Создайте Secret для чувствительных данных:

```yaml
apiVersion: v1
kind: Secret
metadata:
  name: newsletter-secret
type: Opaque
stringData:
  CLIENT_TOKEN: "your-secret-token"
  DB_USER: "newsletter_user"
  DB_PASSWORD: "newsletter_password"
```

Задеплойте приложение:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: newsletter-backend
spec:
  replicas: 3
  selector:
    matchLabels:
      app: newsletter-backend
  template:
    metadata:
      labels:
        app: newsletter-backend
    spec:
      containers:
      - name: newsletter-backend
        image: newsletter-backend:latest
        ports:
        - containerPort: 8080
        envFrom:
        - configMapRef:
            name: newsletter-config
        - secretRef:
            name: newsletter-secret
```

## Рекомендации по безопасности

1. **Используйте сложный CLIENT_TOKEN** — генерируйте длинные случайные токены
2. **Используйте HTTPS в продакшене** — размещайте за обратним прокси (nginx, traefik)
3. **Регулярно меняйте пароли и токены**
4. **Минимизируйте права пользователя БД**
5. **Включите ограничение частоты запросов** — добавьте middleware rate limiting
6. **Мониторьте логи** — настройте централизованный сбор логов (ELK, Loki и др.)

## Тестирование

### Ручное тестирование

```bash
# Валидный запрос
curl -X POST http://localhost:8080/api/newsletter/subscribe \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "clientToken": "your-token"}'

# Неверный email
curl -X POST http://localhost:8080/api/newsletter/subscribe \
  -H "Content-Type: application/json" \
  -d '{"email": "invalid-email", "clientToken": "your-token"}'

# Неверный токен
curl -X POST http://localhost:8080/api/newsletter/subscribe \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "clientToken": "wrong-token"}'

# Дублирующийся email
curl -X POST http://localhost:8080/api/newsletter/subscribe \
  -H "Content-Type: application/json" \
  -d '{"email": "test@example.com", "clientToken": "your-token"}'
```

## Решение проблем

### Нет подключения к базе данных

Проверьте, что база данных запущена и доступна. Если вы используете внешний docker-compose репозиторий — проверьте его логи и статус контейнеров.

### Миграции не применились

Запустите приложение повторно или примените миграции вручную с помощью golang-migrate.

### Порт уже занят

Измените порт в переменных окружения или при запуске контейнера:
```bash
docker run -d -p 8081:8080 ...
```

## Лицензия

MIT

## Поддержка

По вопросам и ошибкам создайте issue в репозитории.
