FROM golang:1.18-alpine

COPY go.mod .
COPY go.sum .
COPY main.go .