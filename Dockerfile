FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY app /app
WORKDIR /
ENTRYPOINT /app
