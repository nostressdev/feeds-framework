FROM golang:1.17 as builder

RUN apt update
RUN apt install -y protobuf-compiler
RUN go install github.com/luckybet100/protodeps@v1.0.3

RUN wget --no-check-certificate https://github.com/apple/foundationdb/releases/download/6.3.22/foundationdb-clients_6.3.22-1_amd64.deb
RUN dpkg -i foundationdb-clients_6.3.22-1_amd64.deb
RUN apt update
RUN apt install -y dnsutils

WORKDIR /service

COPY ./go.mod /service

COPY . /service

RUN make bin/feeds
