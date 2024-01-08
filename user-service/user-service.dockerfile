FROM alpine:latest

WORKDIR /app

COPY /dist/user-service .

CMD [ "./user-service" ]