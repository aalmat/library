FROM golang:1.20
WORKDIR /app

COPY ./ ./

RUN go mod download

RUN go build -o bookstore ./cmd/main.go

EXPOSE 8080

CMD ["./bookstore"]