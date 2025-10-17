FROM golang:1.25.1-alpine

WORKDIR /app

COPY DataBaseService/go.mod DataBaseService/go.sum ./
RUN go mod download

COPY DataBaseService/ ./

RUN go build -o main ./cmd/databaseservice/main.go

EXPOSE 50051

CMD ["./main"]