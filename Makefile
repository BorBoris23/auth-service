include .env
export

migrate-up:
	migrate -path migrations -database "$(DATABASE_URL)" up

migrate-down:
	migrate -path migrations -database "$(DATABASE_URL)" down 1

migrate-version:
	migrate -path migrations -database "$(DATABASE_URL)" version

seed:
	go run cmd/seed/main.go

auth:
	go run cmd/auth/main.go

docker-up:
	docker compose up -d

docker-down:
	docker compose down