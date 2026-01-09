# syntax=docker/dockerfile:1
ARG GO_VERSION=1.22
FROM golang:${GO_VERSION}-alpine AS build

WORKDIR /src
COPY go.mod ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /out/devctl .

FROM alpine:3.20
RUN adduser -D -g '' devctl
USER devctl
COPY --from=build /out/devctl /usr/local/bin/devctl

ENTRYPOINT ["devctl"]
