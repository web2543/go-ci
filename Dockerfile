FROM alpine:latest
COPY api .
EXPOSE 3000/tcp
ENTRYPOINT [ "./","api" ]