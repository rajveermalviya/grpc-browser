FROM envoyproxy/envoy:latest
COPY ./envoy.yaml /etc/envoy/envoy.yaml
CMD /usr/local/bin/envoy -c /etc/envoy/envoy.yaml

# docker build -t grpc/envoy -f ./envoy.Dockerfile .
# docker run --rm -d -p 8080:8080 --network=host grpc/envoy