# Dockerfile
FROM alpine
COPY jbot /usr/bin/jbot
CMD ["/usr/bin/jbot"]
