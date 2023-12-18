FROM golang:1.21.5 AS builder

WORKDIR /build

COPY ./src ./

RUN go build -o /go-binary-app

FROM debian:12 as runner

WORKDIR /app

COPY --from=builder /go-binary-app ./

ENTRYPOINT ["/app/go-binary-app"]