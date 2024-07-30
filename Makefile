build:
	docker compose -f docker-compose.yaml build --no-cache
up:
	docker compose -f docker-compose.yaml up -d
down:
	docker compose -f docker-compose.yaml down
logs:
	docker logs --since 120s web -f
exec:
	docker exec -it web sh
mysql:
	docker compose -f docker-compose.yaml exec db mysql -u root -p
