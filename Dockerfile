FROM alpine:3.17
WORKDIR /work
COPY --chmod=755 api .
EXPOSE 3000/tcp
ENV GIN_MODE=release
CMD ["./api"]
