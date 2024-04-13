FROM golang as builder

WORKDIR /root/go/
COPY . .
RUN go get .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o xray-iran-bridge-updater .


FROM alpine:3.17
#RUN apk --no-cache add ca-certificates

RUN wget https://github.com/XTLS/Xray-core/releases/download/v1.8.10/Xray-linux-64.zip
RUN unzip Xray-linux-64.zip -d /root/xray-core

WORKDIR /root/xray-core

RUN wget https://github.com/bootmortis/iran-hosted-domains/releases/latest/download/iran.dat
RUN wget https://github.com/MrMohebi/domain-list-iran-bans/releases/latest/download/iran-bans.dat

COPY ./configs /root/xray-core/configs
COPY --chmod=777 ./rebootXray.sh /root/xray-core/rebootXray.sh
COPY --from=builder --chmod=777 /root/go/xray-iran-bridge-updater /root/xray-core/xray-iran-bridge-updater


#CMD ["/root/xray-core/xray", "run", "-confdir", "/root/xray-core/configs"]
