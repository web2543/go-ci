FROM alpine:latest
COPY . .
EXPOSE 3000/tcp
ENTRYPOINT [ "go","run","main.go" ]