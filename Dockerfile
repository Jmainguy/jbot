# Dockerfile
FROM alpine:3.18.5
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
