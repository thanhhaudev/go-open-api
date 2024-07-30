build:
	docker compose -f docker-compose.yaml build --nocache
up:
	docker compose -f docker-compose.yaml up -d
down:
	docker compose -f docker-compose.yaml down
logs:
	docker logs --since 120s web -f
