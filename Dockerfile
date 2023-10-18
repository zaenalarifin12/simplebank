# Build Stage
FROM golang:1.21.3-alpine3.18 AS builder
WORKDIR /app
COPY . .
RUN go build -o main .

# Run Stage
FROM alpine:3.18
WORKDIR /app
COPY --from=builder /app .
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./db/migration

RUN chmod +x /app/start.sh
RUN chmod +x /app/wait-for.sh

EXPOSE 8080
CMD ["/app"]
