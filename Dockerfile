# syntax=docker/dockerfile:1

###
### BUILD
###
#FROM golang:1.16 AS builder
#WORKDIR /build
#
#COPY go.mod ./
#COPY go.sum ./
#RUN go mod download
#
#COPY . ./
##RUN GOOS=linux GOARCH=amd64 go build -a -ldflags '-s' .
#RUN make deploy

###
### DEPLOY
###
FROM alpine:3.14
#FROM scratch
LABEL maintainer="gainchang620@gmail.com"
WORKDIR /app

ENV GIN_MODE=release
#ADD ./config ./config
#COPY ./config ./config

#COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
#COPY --from=builder /etc/passwd /etc/passwd
#COPY --from=builder /build/conctl .

EXPOSE 8000
#ENTRYPOINT [""]