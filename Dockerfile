# Dockerfile
FROM alpine:3
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
