# Newsletter service

Этот сервис предназначен для сайта музыкальной группы Infernus Presence.

Готовый к продакшену Go-бэкенд для подписки на рассылку с использованием базы данных MariaDB.

## Возможности

- ✅ RESTful API для подписки на рассылку
- ✅ Валидация email-адреса
- ✅ Аутентификация клиента по токену
- ✅ MariaDB с миграциями
- ✅ Поддержка Docker
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
- **Контейнеризация**: Docker и Docker Compose

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
├── docker-compose.yml             # Конфигурация Docker Compose
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

### Требования

- Docker
- Docker Compose

### Шаг 1: Клонируйте или распакуйте проект

```bash
cd newsletter-backend
```

### Шаг 2: Настройте окружение

Отредактируйте `docker-compose.yml` и задайте `CLIENT_TOKEN`:

```yaml
environment:
  CLIENT_TOKEN: "your-secret-client-token-here"  # Измените это!
```

### Шаг 3: Запустите приложение

```bash
docker-compose up -d
```

Это:
- Запустит контейнер MariaDB
- Соберёт и запустит Go-приложение
- Автоматически применит миграции
- Откроет API на `http://localhost:8080`

### Шаг 4: Проверьте API

```bash
curl -X POST http://localhost:8080/api/newsletter/subscribe \
  -H "Content-Type: application/json" \
  -d '{
    "email": "test@example.com",
    "clientToken": "your-secret-client-token-here"
  }'
```

### Шаг 5: Просмотр логов

```bash
docker-compose logs -f app
```

### Остановить приложение

```bash
docker-compose down
```

Для удаления томов:

```bash
docker-compose down -v
```

## Локальная разработка (без Docker)

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

# Запуск контейнеров Docker
make docker-up

# Остановка контейнеров Docker
make docker-down

# Просмотр логов Docker
make docker-logs
```

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

Проверьте, что база данных запущена:
```bash
docker-compose ps
```

Проверьте логи базы данных:
```bash
docker-compose logs db
```

### Миграции не применились

Ручной запуск миграций:
```bash
docker-compose exec app ./main
```

Или подключитесь к базе и выполните SQL вручную:
```bash
docker-compose exec db mysql -u newsletter_user -p newsletter
```

### Порт уже занят

Измените порт в `docker-compose.yml`:
```yaml
ports:
  - "8081:8080"  # Используйте другой порт хоста
```

## Лицензия

MIT

## Поддержка

По вопросам и ошибкам создайте issue в репозитории.
