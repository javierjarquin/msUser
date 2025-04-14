FROM golang:1.23 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Copiamos .env explícitamente, por si está ignorado
COPY .env .env

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app .

FROM alpine:latest
WORKDIR /root/
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/app .
COPY --from=builder /app/.env .env  
EXPOSE 8081
CMD ["./app"]
