FROM golang:1.20.5-alpine3.17 as builder

WORKDIR /app
COPY . .

RUN apk add --no-cache \
  ca-certificates \
  curl \
  git

RUN CGO_ENABLED=0 \
  go build -v -o app

FROM alpine:3.18.2
WORKDIR /app
RUN apk add --no-cache ca-certificates
COPY --from=builder /app/app /app/app
RUN chmod +x app
CMD ["/app/app"]
