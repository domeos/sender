FROM centos:7

MAINTAINER domeos

# Prepare

RUN rm -f /etc/localtime && ln -s /usr/share/zoneinfo/PRC /etc/localtime && mkdir -p /opt/falcon-sender && cd /opt/falcon-sender

# Set work-dir

WORKDIR /opt/falcon-sender

# Copy falcon-sender files

COPY falcon-sender ./falcon-sender
COPY control ./control
COPY cfg.template.json ./cfg.template.json

# Copy start script

COPY run.sh ./run.sh

# Copy dockerize

COPY dockerize /usr/local/bin/dockerize

# Start service

CMD dockerize -template ./cfg.template.json:./cfg.json sh run.sh
