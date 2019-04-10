FROM alpine:latest

COPY ./build .

EXPOSE 8888

ENTRYPOINT [ "./app" ] 