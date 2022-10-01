postgres:
	sudo docker run --name pgdatabase -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -p 8081:5432 -d postgres:12-alpine

createdb:
	sudo docker exec -it pgdatabase createdb --username=root --owner=root pathshaala

dropdb:
	sudo docker exec -it postgres12 dropdb pathshaala

migrateup:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/pathshaala?sslmode=disable" -verbose up 

migratedown:
	migrate -path db/migration -database "postgresql://root:root@localhost:8081/pathshaala?sslmode=disable" -verbose down

live:
	sudo docker start pgdatabase ; nodemon

.PHONY: postgres createdb dropdb migrateup migratedown