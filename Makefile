dev-up:
	@docker compose \
		-f docker/docker-compose.yml up -d

dev-down:
	@docker compose \
		-f docker/docker-compose.yml down

migrate-up:
	@migrate -path migration -database "postgres://root:password@localhost:5432/pairswap?sslmode=disable" -verbose up

migrate-down:
	@migrate -path migration -database "postgres://root:password@localhost:5432/pairswap?sslmode=disable" -verbose down
