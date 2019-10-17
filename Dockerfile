FROM golang:1.12 as builder
LABEL MANTAINER="Sergei Bukharkin <serboox@gmail.com>"

# If you wanna remake to Alpine before install show packages
# RUN apk add --no-cache git bash make ca-certificates curl gcc musl-dev

COPY . /go/src/github.com/serboox/os-cli
WORKDIR /go/src/github.com/serboox/os-cli
RUN make build-for-docker

FROM ubuntu:18.04
RUN apt update -y &&\
    apt install -y ca-certificates curl &&\
    rm -rf /var/lib/apt/lists/*
COPY --from=builder /go/src/github.com/serboox/os-cli ./
