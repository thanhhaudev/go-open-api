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
