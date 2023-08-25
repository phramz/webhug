# builder: golang
FROM golang:1.21-alpine as builder

WORKDIR /.build

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

# see https://medium.com/@diogok/on-golang-static-binaries-cross-compiling-and-plugins-1aed33499671
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -tags netgo -ldflags '-w -extldflags "-static"' -o ./bin/webhug cmd/webhug.go

# final layer
FROM alpine:latest as final

COPY --from=alpine:latest /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /.build/bin/webhug /bin/webhug
COPY --from=builder /.build/config.yaml /etc/webhug/config-example.yaml

ARG BUILD_COMMIT=dev
ENV BUILD_COMMIT=$BUILD_COMMIT
ARG BUILD_VERSION=dev
ENV BUILD_VERSION=$BUILD_VERSION

ENTRYPOINT ["/bin/webhug"]
