FROM golang:1.23.4-alpine3.20 AS builder

RUN apk update && \
    apk add --no-cache \
    ffmpeg \
    ca-certificates \
    build-base

WORKDIR /var/app/videoverse/

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN make build  #build the project

FROM alpine:3.20

WORKDIR /var/app/videoverse/

RUN apk add --no-cache bash jq \
    && apk update && apk upgrade

COPY --from=builder /var/app/videoverse/bin bin/

EXPOSE 80

CMD [ "./bin/videoverse" ]
