FROM alpine:3.17
RUN wget https://github.com/XTLS/Xray-core/releases/download/v1.8.3/Xray-linux-64.zip
RUN unzip Xray-linux-64.zip -d /root/xray-core

WORKDIR /root/xray-core

RUN wget https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/iran.dat
RUN wget https://github.com/MrMohebi/domain-list-iran-bans/releases/latest/download/iran-bans.dat

COPY ./configs /root/xray-core/configs

CMD ["/root/xray-core/xray", "run", "-confdir", "/root/xray-core/configs"]
