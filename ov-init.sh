#!/bin/bash -x

# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs -d '\n')

MY_IP=$(curl ifconfig.io)

docker run -v "$OV_VOLUME":/etc/openvpn --rm docker.iranrepo.ir/kylemanna/openvpn ovpn_genconfig -u udp://"$MY_IP":"$OV_PORT"
docker run -d --rm --name dummy -v "$OV_VOLUME":/root alpine tail -f /dev/null
docker cp ./my_ovpn_initpki.sh dummy:/root/my_ovpn_initpki.sh
docker stop dummy
docker run -v "$OV_VOLUME":/etc/openvpn --rm -it docker.iranrepo.ir/kylemanna/openvpn /etc/openvpn/my_ovpn_initpki.sh nopass
