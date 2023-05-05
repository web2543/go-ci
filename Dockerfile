FROM alpine:3.17
WORKDIR /work
COPY --chmod=755 api .
EXPOSE 3000/tcp
ENTRYPOINT ["pwd"]
