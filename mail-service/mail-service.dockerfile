FROM alpine:latest

WORKDIR /app

COPY /dist/mail-service .

CMD [ "./mail-service" ]