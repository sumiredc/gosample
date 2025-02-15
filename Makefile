.PHONY: mysql migrate

mysql: 
	docker compose exec mysql mysql -u user -h localhost -ppassword sample

migrate:
	docker compose run --rm migration up
