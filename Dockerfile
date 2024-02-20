FROM golang:1.22-alpine3.19 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY scheduler.env .
COPY schedules.yml .

EXPOSE 8082
CMD ["/app/main"]