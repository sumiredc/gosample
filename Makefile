.PHONY: mysql migrate, test, coverage

mysql: 
	docker compose exec mysql mysql -u user -h localhost -ppassword sample

migrate:
	docker compose run --rm migration up

test:
	docker compose run --rm go go test ./...

coverage:
	docker compose run --rm go go test -cover ./...

coverage-detail:
	docker compose run --rm go bash -c "go test ./... -covermode=atomic -coverprofile=coverage.out && go tool cover -func=coverage.out > coverage.atomic.out"
