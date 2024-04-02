FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum /
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -v ./cmd/shop

EXPOSE 3000

ENTRYPOINT ["./shop"]
