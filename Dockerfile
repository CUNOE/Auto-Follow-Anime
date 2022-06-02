FROM golang:1.17.8 as builder

MAINTAINER cunoe

ENV GO111MODULE=on \
    CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /go/src/autoFollowAnime

COPY . .

RUN  go build -o app

FROM ubuntu:20.04 as prod

WORKDIR /root/

COPY --from=builder /go/src/autoFollowAnime/app .

RUN apt-get update && apt-get -qq install -y --no-install-recommends ca-certificates curl

CMD ["./app"]
