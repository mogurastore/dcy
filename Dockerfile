FROM golang:1.18 AS builder
WORKDIR /tmp/gobuild
COPY . ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

FROM scratch
COPY --from=builder /tmp/gobuild/dcy /bin/dcy
ENTRYPOINT ["/bin/dcy"]
