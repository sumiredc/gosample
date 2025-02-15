FROM golang:1.24.0

WORKDIR /app

COPY go.mod .

RUN go mod tidy

CMD ["go", "run", "user/main.go"]
