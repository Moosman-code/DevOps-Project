FROM golang:1.20-alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY main.go .

RUN go build -o app main.go

FROM alpine:latest
WORKDIR /root/

COPY --from=builder /app/app .

EXPOSE 80

CMD ["./app"]
