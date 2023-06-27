#!/bin/bash

# shellcheck disable=SC2046
export $(grep -v '^#' .env | xargs -d '\n')

# Check if two arguments are not entered
if [ $# -lt 2 ]
then
  echo "Error: two arguments are required (CLIENT_NAME and CLIENT_PASS)"
  exit 1
fi

CLIENT_NAME=$OV_PREFIX$1
CLIENT_PASS=$2

docker run -v "$OV_VOLUME":/etc/openvpn --rm -it docker.iranrepo.ir/kylemanna/openvpn /bin/bash -c "EASYRSA_PASSOUT=pass:$CLIENT_PASS EASYRSA_PASSIN=pass:$CLIENT_PASS easyrsa build-client-full $CLIENT_NAME"
docker run -v "$OV_VOLUME":/etc/openvpn --rm -it docker.iranrepo.ir/kylemanna/openvpn ovpn_getclient "$CLIENT_NAME" > "$CLIENT_NAME".ovpn


docker run -v ./"$CLIENT_NAME".ovpn:/root/repo/"$CLIENT_NAME".ovpn --rm -it docker.iranrepo.ir/alpine sh -c "
apk update && apk add git && \
git clone https://mrm:$TOKEN_GITHUB@github.com/$REPO repo  && \
cd repo && \
git config --local user.name auto-commit && \
git config --local user.email "auto-commit@users.noreply.github.com" && \
git checkout -b clients && \
git rm -rf $(ls | grep -v *.ovpn) && \
cp /root/repo/$CLIENT_NAME.ovpn ./ && \
git add $CLIENT_NAME.ovpn && \
git commit -m 'add new client => $CLIENT_NAME with password => $CLIENT_PASS' && \
git push -f -u origin clients
"


