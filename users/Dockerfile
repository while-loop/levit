FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git

RUN go get github.com/while-loop/levit
RUN cd ${GOPATH}/src/github.com/while-loop/levit
RUN go build -o usersd ./users/

FROM busybox:latest
COPY --from=builder /go/bin/usersd /usr/local/bin/
RUN usersd version

CMD ["usersd"]