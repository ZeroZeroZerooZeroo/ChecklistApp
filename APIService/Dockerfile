FROM golang:1.25.1-alpine

WORKDIR /app

COPY APIService/go.mod APIService/go.sum ./
RUN go mod download

COPY APIService/ ./

RUN go build -o main ./cmd/apiservice/main.go

EXPOSE 8080

CMD ["./main"]