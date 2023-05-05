FROM debian:bullseye
WORKDIR /work
COPY --chmod=755 api .
EXPOSE 3000/tcp
ENTRYPOINT [ "sh","./api"]