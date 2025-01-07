FROM ubuntu:24.04

RUN apt update
RUN apt install -y ffmpeg supervisor htop vim

WORKDIR /media
RUN mkdir -p /var/log/supervisor

COPY supervisord.conf /etc/supervisor/conf.d/supervisord.conf

CMD ["/usr/bin/supervisord"]