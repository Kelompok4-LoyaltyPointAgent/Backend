# build stage
FROM golang:1.19-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main.app 

# final stage
FROM alpine:latest
COPY --from=builder /app/main.app /app/main.app
COPY --from=builder /app/html /app/html
WORKDIR /app
ENTRYPOINT ["./main.app"]
LABEL Name=loyalty-point-agent Version=1.0
