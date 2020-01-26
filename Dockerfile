FROM golang:1.13-alpine3.11 as base

RUN mkdir -p /code/build
WORKDIR /code

COPY . .
RUN apk add --no-cache build-base && make build

FROM docker:19.03 as final

ARG RELEASE_VERSION
ENV RELEASE_VERSION=${RELEASE_VERSION}

RUN mkdir -p /root/.webhug/
WORKDIR /etc/webhug/
COPY --from=base /code/webhug /usr/bin/webhug
COPY --from=base /code/config.yaml /root/.webhug/config.yaml

EXPOSE 8080

ENTRYPOINT ["/usr/bin/webhug"]

