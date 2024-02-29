# Dockerfile
FROM alpine:3.18.6
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
