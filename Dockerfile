FROM golang:1.11 as build

ADD . /go/src/github.com/dazwilkin/metadata
WORKDIR /go/src/github.com/dazwilkin/metadata

ARG GOFLAGS=""
RUN go get ${GOFLAGS} ./...

FROM gcr.io/distroless/base

COPY --from=build /go/bin/metadata /

ENTRYPOINT ["/metadata"]
