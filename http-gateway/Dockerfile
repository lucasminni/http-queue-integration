FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . ./

RUN go mod download

COPY *.go ./

RUN go build -o http-gateway-app .

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/http-gateway-app .

RUN apk add curl

EXPOSE 80

CMD ["./http-gateway-app"]
