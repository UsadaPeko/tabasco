FROM golang:1.18.1-alpine as Builder

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build /app/cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main ./main

ENTRYPOINT ["./main"]
