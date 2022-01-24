FROM golang:1.16-alpine as builder

WORKDIR /go/src/url

COPY . .

RUN go get -d -v ./...

RUN go build -o /app/url ./cmd/url/.

FROM alpine:3.15

COPY --from=builder /app/url /app/url

COPY configs/config.yaml /app/config.yaml

WORKDIR /app

ENTRYPOINT ["/app/url"]