FROM golang:1.19 as builder

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
    && apt-get upgrade -y

WORKDIR /go/src

COPY ./go.mod .
COPY ./go.sum .
RUN go mod download

COPY ./cmd/ ./cmd/
COPY ./internal/ ./internal/
COPY ./pkg/ ./pkg/
RUN GOOS=linux GOARCH=amd64 go build -ldflags "-w -s" -o /go/bin/webapi-api ./cmd/api/main.go

FROM ubuntu:22.04 as runner

ENV DEBIAN_FRONTEND=noninteractive
RUN apt-get update \
  && apt-get upgrade -y \
  && apt-get install -y ca-certificates tzdata \
  && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

COPY --from=builder /go/bin/webapi-api /opt/webapi/bin/webapi-api

EXPOSE 4000

CMD ["/opt/webapi/bin/webapi-api"]
