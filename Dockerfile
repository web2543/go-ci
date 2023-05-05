FROM alpine:3.17
COPY --chmod=755 api .
EXPOSE 3000/tcp
ENTRYPOINT [ "./","api"]