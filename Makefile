include .env

migrateup:
	migrate -path db/migrations -database "${DB_TYPE}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "${DB_TYPE}://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=disable" -verbose down $N

mock-generate: 	
	mockgen -package mock -destination repository/mock/store.go github.com/bobbybof/timesheet/internal/repository Store

sqlc:
	sqlc generate && make mock-generate

seed:
	 psql "postgresql://${DB_USERNAME}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}" -f ./db/seed.sql

test:
	go test -v -cover ./...

server:
	go run main.go 

.PHONY: server migrateup migratedown sqlc