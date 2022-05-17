FROM alpine
EXPOSE 50051
COPY grpc-qs grpc-qs
COPY . .
COPY grpc_health_probe-linux-amd64 /bin/
RUN chmod +x /bin/grpc_health_probe-linux-amd64
ENTRYPOINT ["/grpc-qs"]