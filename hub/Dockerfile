FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git make

RUN go get github.com/while-loop/levit/hub/...
WORKDIR ${GOPATH}/src/github.com/while-loop/levit/hub
RUN make all

FROM alpine:latest
COPY --from=builder /go/bin/hubd /usr/local/bin/
RUN hubd -v

ENTRYPOINT ["hubd"]