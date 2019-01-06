FROM golang:alpine AS builder
RUN apk update && apk upgrade && apk add --no-cache git ca-certificates gcc musl-dev
ENV GO111MODULE=on
WORKDIR /src/app
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go fix ./...
RUN go build -o /app

FROM alpine
RUN apk update && apk upgrade && apk add --no-cache bash ca-certificates curl htop
COPY --from=builder /app /app
ENTRYPOINT /app
