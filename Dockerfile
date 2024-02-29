# Dockerfile
FROM alpine:3.19.1
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
