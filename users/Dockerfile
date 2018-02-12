FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk add --update --no-cache git make build-base

RUN go get github.com/while-loop/levit/users/...
WORKDIR ${GOPATH}/src/github.com/while-loop/levit/users
RUN make all

FROM alpine:latest
COPY --from=builder /go/bin/usersd /usr/local/bin/
RUN usersd -v

ENTRYPOINT ["usersd"]