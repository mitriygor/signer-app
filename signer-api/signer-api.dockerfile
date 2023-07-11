FROM alpine:latest

RUN echo "net.ipv4.ip_local_port_range = 1024 65535" > /etc/sysctl.conf && \
    echo "net.ipv4.tcp_fin_timeout = 30" >> /etc/sysctl.conf && \
    echo "net.ipv4.tcp_tw_reuse = 1" >> /etc/sysctl.conf

RUN mkdir /app

COPY signerApi /app

CMD [ "/app/signerApi"]