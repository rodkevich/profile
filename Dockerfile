FROM golang:1.15.2-alpine3.12

RUN apk add --no-cache \
        libc6-compat

WORKDIR /go/src
COPY . profile

RUN chmod -R 777 profile
RUN chmod +X profile
CMD /go/src/profile
