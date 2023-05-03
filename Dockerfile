FROM alpine:3.17
COPY api .
EXPOSE 3000/tcp
ENTRYPOINT [ "./","api"]