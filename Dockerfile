# Dockerfile
FROM alpine:latest
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
