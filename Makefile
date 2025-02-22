.PHONY: mysql migrate, test, coverage

mysql: 
	docker compose exec mysql mysql -u user -h localhost -ppassword sample

migrate:
	docker compose run --rm migration up

test:
	docker compose run --rm go go test ./...

coverage:
	docker compose run --rm go go test -cover ./...
