# build
FROM golang:1.12-alpine as builder
COPY            . /go/src/github.com/moul/grpcbin
WORKDIR         /go/src/github.com/moul/grpcbin
RUN             GO111MODULE=off go build -o /go/bin/grpcbin -ldflags "-extldflags \"-static\"" -v

# minimal runtime
FROM            alpine
RUN             apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
COPY            --from=builder /go/bin/grpcbin /bin/grpcbin
COPY            --from=builder /go/src/github.com/moul/grpcbin/cert /root/cert
WORKDIR         /root
EXPOSE          9000 9001 80
ENTRYPOINT      ["/bin/grpcbin"]
