start-db:
	docker start ecoforums-db
stop-db:
	docker stop ecoforums-db
run-postgres:
	docker exec -it ecoforums-db psql -U postgres

create-table:
	migrate create -ext sql -dir db/migrations -seq create_$(name)_table
db-up:
	migrate -database "postgresql://postgres:postgres@localhost:5436/ecoforums?sslmode=disable" -path db/migrations --verbose up

db-down:
	migrate -database "postgresql://postgres:postgres@localhost:5436/ecoforums?sslmode=disable" -path db/migrations --verbose down