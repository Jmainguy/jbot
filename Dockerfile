# Dockerfile
FROM alpine:3.18.4
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
