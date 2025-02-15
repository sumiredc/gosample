FROM golang:1.24.0

WORKDIR /app/migrations

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@v4.18.1

ENTRYPOINT [ "migrate", "-database", "mysql://user:password@tcp(mysql:3306)/sample", "-source", "file:///app/migrations" ]
