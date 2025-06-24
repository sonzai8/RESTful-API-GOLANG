include .env
export
MIGRATE_DIR = internal/db/migrations

#DB_HOST=localhost
#DB_USER=root
#DB_PASSWORD=sonzai@123456
#DB_NAME=master-golang
#DB_PORT=5432
#DB_SSLMODE=disable

CONN_STRING = postgres://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=$(DB_SSLMODE)

importdb:
	docker exec -i postgres-db psql -U root -d master-golang < ./backupdb-master-golang.sql
exportdb:
	docker exec -i postgres-db pg_dump -U root -d master-golang > ./backupdb-master-golang.sql
dev:
	go run .
# create a new migration (make migrate-create name=profiles)
migrate-create:
	migrate create -ext sql -dir $(MIGRATE_DIR) -seq $(name)
#Run all pening migration (make migrate-up)
migrate-up:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) up
# apply specific migration version make migrate-togo version=1
migrate-goto:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) goto $(version)

# Rollback the last migration
migrate-down:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) down 1
# Rollback n migrations
migrate-down-n:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) down $(n)

# force migrate version (user with caution example : make migrate-force version=1)
migrate-force:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) force $(version)
#
# Drop everything (include schema migration)
migrate-drop:
	migrate -path $(MIGRATE_DIR) -database $(CONN_STRING) drop

.PHONY: importdb exportdb dev migrate-up migrate-down migrate-create migrate-force migrate-drop migrate-goto migrate-down-n