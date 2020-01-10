FROM golang:1.13 AS builder
ADD . /tmp/service-build
WORKDIR /tmp/service-build
RUN make build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /tmp/service-build/example-service .
CMD ["./example-service", "start"]