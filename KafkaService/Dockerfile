FROM golang:1.25.1-alpine

WORKDIR /app

COPY KafkaService/go.mod KafkaService/go.sum ./
RUN go mod download

COPY KafkaService/ ./

RUN go build -o main ./cmd/kafkaservice/main.go

CMD ["./main"]