FROM golang:1.21.5-alpine

WORKDIR /app

# COPY go.mod go.sum ./
# RUN go mod download

COPY . .

COPY migration.sql ./

RUN go build -o ./main cmd/main.go

ENTRYPOINT [ "./main" ]