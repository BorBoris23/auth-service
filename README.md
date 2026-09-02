Auth Service

Сервис авторизации на Go + PostgreSQL.

Docker
Запустить PostgreSQL
make docker-up

или:

docker compose up -d
Проверить запущенные контейнеры
docker ps
Остановить PostgreSQL
make docker-down

или:

docker compose down
Остановить PostgreSQL и удалить volumes

⚠️ Команда удалит данные PostgreSQL.

docker compose down -v
Migrations
Накатить все миграции
make migrate-up
Откатить последнюю миграцию
make migrate-down
Посмотреть текущую версию миграций
make migrate-version
Создать новую миграцию
migrate create -ext sql -dir migrations -seq migration_name

После выполнения команды появятся:

000XXX_migration_name.up.sql
000XXX_migration_name.down.sql
Seeds

Сиды добавляют тестовые данные в базу.

Запустить сиды
make seed

Сидер добавляет:

5 ролей;
10 пользователей;
одного пользователя с ролью admin;
остальных пользователей с другими ролями;
пароли пользователей сохраняются в виде bcrypt-хэшей.

Сиды можно запускать повторно благодаря ON CONFLICT DO NOTHING.

Environment

В корне проекта находится .env:

DATABASE_URL=postgres://auth:auth@localhost:5432/auth?sslmode=disable

JWT_SECRET=your-secret-key

Переменные окружения автоматически загружаются приложением через godotenv.

Выполнять source .env не требуется.

Project Structure
auth-service/
│
├── cmd/
│   ├── auth/
│   │   └── main.go
│   │
│   └── seed/
│       └── main.go
│
├── internal/
│   ├── auth/
│   │   └── auth.go
│   │
│   ├── http/
│   │   ├── auth.go
│   │   ├── dto/
│   │   │   └── login.go
│   │   ├── response.go
│   │   └── router.go
│   │
│   ├── jwt/
│   │   └── jwt.go
│   │
│   ├── postgres/
│   │   └── postgres.go
│   │
│   ├── seeds/
│   │   ├── roles.go
│   │   └── users.go
│   │
│   └── user/
│       └── user.go
│
├── migrations/
│   ├── 000001_create_roles.up.sql
│   ├── 000001_create_roles.down.sql
│   ├── 000002_create_users.up.sql
│   └── 000002_create_users.down.sql
│
├── docker-compose.yml
├── Makefile
├── go.mod
└── go.sum
Authentication

Авторизация выполняется через:

POST /login
Процесс авторизации
HTTP Request
     ↓
LoginRequest
     ↓
Validation
     ↓
FindUser
     ↓
Check password with bcrypt
     ↓
FindRoleByID
     ↓
Generate JWT
     ↓
LoginResponse
Успешная авторизация

Сервис возвращает HTTP 200 OK:

{
  "message": "login successful",
  "token": "JWT_TOKEN",
  "user": {
    "id": 1,
    "login": "admin",
    "role": "admin"
  }
}
Неверный логин или пароль

Сервис возвращает HTTP 401 Unauthorized:

{
  "error": "invalid_credentials",
  "message": "Invalid login or password"
}
Основные команды
Запустить PostgreSQL
make docker-up
Остановить PostgreSQL
make docker-down
Накатить миграции
make migrate-up
Откатить последнюю миграцию
make migrate-down
Проверить версию миграций
make migrate-version
Запустить сиды
make seed
Запустить auth-service
make auth