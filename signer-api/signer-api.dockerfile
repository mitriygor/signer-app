FROM alpine:latest

RUN mkdir /app

COPY signerApi /app

CMD [ "/app/signerApi"]