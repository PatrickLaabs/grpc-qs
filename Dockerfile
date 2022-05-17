FROM alpine
EXPOSE 8080
COPY restfulapi /usr/bin/restfulapi
ENTRYPOINT ["/usr/bin/restfulapi"]