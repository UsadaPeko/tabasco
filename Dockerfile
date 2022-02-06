FROM golang:1.17.6-alpine as Builder

RUN mkdir -p /app

WORKDIR /app

COPY . .

RUN go mod download

RUN go build /app/cmd/main.go

FROM alpine:latest

COPY --from=builder /app/main ./main

ENTRYPOINT ["./main"]
