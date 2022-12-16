migrateup:
	migrate -path db/migration -database "postgresql://postgres:3drTCFvaKD@localhost:5432/root?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:3drTCFvaKD@localhost:5432/root?sslmode=disable" -verbose down

run:
	go run server.go