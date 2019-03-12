FROM golang AS builder
WORKDIR /src/app
COPY ./ ./
RUN go build -mod=vendor -o /app

FROM debian
RUN apt update && DEBIAN_FRONTEND=noninteractive apt upgrade -y && DEBIAN_FRONTEND=noninteractive apt install ca-certificates curl htop -y
COPY --from=builder /app /app
ENTRYPOINT /app
EXPOSE 9000