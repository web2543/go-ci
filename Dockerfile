FROM golang:1.20.3-alpine
COPY --chmod=755 api .
EXPOSE 3000/tcp
ENV GIN_MODE=release
CMD ["/api"]
