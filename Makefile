migrateup:
	migrate -path db/migration -database "postgresql://postgres:3drTCFvaKD@localhost:5432/dotask?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:3drTCFvaKD@localhost:5432/dotask?sslmode=disable" -verbose down

run:
	go run server.go

postgresup:
	docker compose up dotask-db -d

postgresdown:
	docker compose down