FROM alpine:latest

RUN mkdir /app

COPY brokerService /app

CMD [ "/app/brokerService"]