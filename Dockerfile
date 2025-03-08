FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init -g main.go
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/main .
COPY .env .
EXPOSE 8080
CMD [ "./main" ]

