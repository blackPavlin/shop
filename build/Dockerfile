FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v ./cmd/shop

EXPOSE 3000

ENTRYPOINT ["./shop"]
