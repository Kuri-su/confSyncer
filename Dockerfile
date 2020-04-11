# step 0
FROM golang:1.14.2-alpine3.11

USER root
LABEL maintainer="misaki.zhcy@gmail.com"

RUN apk update && apk add go git musl-dev xz binutils \
    && export GO111MODULE=on \
    && export GOPATH=/root/go \
    && go get github.com/Kuri-su/confSyncer \
    && go install github.com/Kuri-su/confSyncer

RUN wget https://github.com/upx/upx/releases/download/v3.96/upx-3.96-amd64_linux.tar.xz \
    && xz -d upx-3.96-amd64_linux.tar.xz \
    && tar -xvf upx-3.96-amd64_linux.tar \
    && cd upx-3.96-amd64_linux \
    && chmod a+x ./upx \
    && mv ./upx /usr/local/bin/ \
    && cd /root/go/bin \
    && strip --strip-unneeded confSyncer \
    && upx confSyncer \
    && chmod a+x ./confSyncer \
    && cp confSyncer /usr/local/bin

# step 1
FROM ubuntu:latest

COPY --from=0 /usr/local/bin/confSyncer /usr/local/bin/

CMD ["confSyncer"," deamon"]