FROM golang:alpine AS builder
ENV GO111MODULE=on
ENV CGO_ENABLED=0
WORKDIR /src/app
RUN apk add --no-cache git
COPY ./go.mod ./go.sum ./
RUN go mod download
COPY ./ ./
RUN go build -o /app

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /app /app
ENTRYPOINT /app