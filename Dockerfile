FROM scratch
EXPOSE 50051
COPY grpc-qs grpc-qs
COPY . .
RUN chmod +x grpc_health_probe-linux-amd64
ENTRYPOINT ["/grpc-qs"]