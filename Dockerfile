FROM golang:1.20.3
COPY . .
EXPOSE 3000/tcp
ENTRYPOINT [ "go","run","main.go" ]