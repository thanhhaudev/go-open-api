DSN := 'mysql://root:rootpwd@tcp(localhost:8989)/playground?charset=utf8&parseTime=True&loc=Local'

build:
	docker compose -f docker-compose.yaml build --no-cache
up:
	docker compose -f docker-compose.yaml up -d
down:
	docker compose -f docker-compose.yaml down
restart:
	docker compose -f docker-compose.yaml restart api && make logs
logs:
	docker logs --since 120s go_api -f
exec:
	docker exec -it go_api sh
mysql:
	docker compose -f docker-compose.yaml exec go_api_db mysql -u root -p
swag:
	swag init --parseDependency
migrate/up:
	migrate -path tools/setup/migrations -database $(DSN) up
migrate/down:
	migrate -path tools/setup/migrations -database $(DSN) down
