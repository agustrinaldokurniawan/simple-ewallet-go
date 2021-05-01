postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root  -e POSTGRES_PASSWORD=root -d postgres:13-alpine

createdb:
	docker exec -it postgres13 createdb --username=root --owner=root ewallet

dropdb:
	docker exec -it postgres13 dropdb ewallet

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/ewallet?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:5432/ewallet?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go clean -testcache && go test -v -cover ./... -count=10

.PHONY: postgres createdb dropdb migrateup migratedown test
