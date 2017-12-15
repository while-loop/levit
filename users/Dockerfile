FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git

RUN go get github.com/while-loop/levit/users/cmd/...

FROM alpine:latest
COPY --from=builder /go/bin/usersd /usr/local/bin/
RUN usersd -v

ENTRYPOINT ["usersd"]