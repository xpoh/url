FROM golang:1.10 AS build
WORKDIR /go/src
COPY ../cmd/url ./go
COPY ../cmd/main.go .

ENV CGO_ENABLED=0
RUN url get -d -v ./...

RUN url build -a -installsuffix cgo -o url .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/url ./
EXPOSE 8080/tcp
ENTRYPOINT ["./url"]
