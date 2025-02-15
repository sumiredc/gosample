# migration

## Create migration files
```sh
docker compose run --rm migration create -ext sql -seq xxx_xxxx_xxxx
```

## Run
```sh
docker compose run --rm migration up {count}
```

## Rollback
```sh
docker compose run --rm migration down

docker compose run --rm migration down {count}
```

## Resolve dirty
```sh
docker compose run --rm migration force {version}
```
