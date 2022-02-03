FROM golang:1.17-bullseye

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.9.0/wait /wait
RUN chmod +x /wait

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

CMD /wait && go run main.go
