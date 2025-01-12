FROM golang:1.23.4-alpine3.20 AS builder

RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    build-base

WORKDIR /var/app/videoverse/

COPY go.mod ./
COPY go.sum ./

COPY . .

RUN make build  #build the project

FROM alpine:3.20

WORKDIR /var/app/videoverse/

RUN apk update && \
    apk add --no-cache \
    ca-certificates \
    ffmpeg

COPY --from=builder /var/app/videoverse/bin .
COPY --from=builder /var/app/videoverse/db ./db

# remove the database file to start with a clean database
RUN rm -f /var/app/videoverse/db/videoverse/videoverse.db

EXPOSE 80

CMD [ "./videoverse" ]