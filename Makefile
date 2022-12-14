build:
	docker compose build bet-app

run:
	docker compose up bet-app -d

migrate:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' up

migrate-down:
	migrate -path ./schema -database 'postgres://postgres:qwerty@0.0.0.0:5432/postgres?sslmode=disable' down