FROM golang:1.17-alpine as build

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY ./app ./app

VOLUME ["./configs/local.config.yml:/app/configs/local.config.yml"]

RUN CGO_ENABLED=0 GOOS=linux go build -v ./app/cmd/topsigner

EXPOSE 3000

CMD ["./shop"]
