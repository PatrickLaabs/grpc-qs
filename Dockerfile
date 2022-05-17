FROM alpine
EXPOSE 50051
COPY grpc-qs grpc-qs
COPY . .
ENTRYPOINT ["/grpc-qs"]