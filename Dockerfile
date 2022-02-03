FROM golang:1.17-bullseye

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD go run main.go
