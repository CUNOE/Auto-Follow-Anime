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

COPY --from=builder /go/src/autoFollowAnime/enter.sh .
COPY --from=builder /go/src/autoFollowAnime/app .

RUN chmod 777 -R /root/enter.sh

CMD ["bash", "-c", "/root/enter.sh && tail -f /dev/null"]