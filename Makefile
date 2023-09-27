postgres:
	docker run --name postgres15-effm -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15-effm createdb --username=root --owner=root simpledb

dropdb:
	docker exec -it postgres15-effm dropdb simpledb

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpledb?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simpledb?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/scipiia/effectivemobiletask/db/sqlc Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock