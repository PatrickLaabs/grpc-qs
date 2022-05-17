FROM scratch
EXPOSE 50051
COPY grpc-qs grpc-qs
COPY . .
ENTRYPOINT ["/grpc-qs"]