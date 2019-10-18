FROM alpine:latest
RUN mkdir /app
RUN mkdir /app/config
WORKDIR /app
ADD ./hello-world /app
ADD ./config/* /app/config/
ENTRYPOINT ["/app/hello-world"]