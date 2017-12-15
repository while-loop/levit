FROM golang:alpine as builder

ENV PATH=${PATH}:${GOPATH}/bin

RUN apk update && apk add git

RUN go get github.com/while-loop/levit/users
RUN cd ${GOPATH}/src/github.com/while-loop/levit/users
RUN go build -o usersd ./...

FROM busybox:latest
COPY --from=builder /go/bin/usersd /usr/local/bin/
RUN usersd -v

CMD ["usersd"]