# Iran bridge Xray config generator
###### **HA:** Multi node support
will generate `outbounds.json` base on provided proxies from [`proxies_active_no_403_under_1000ms.txt`](https://raw.githubusercontent.com/MrMohebi/xray-proxy-grabber-telegram/master/collected-proxies/xray-json/actives_no_403_under_1000ms.txt) 
and update hosted xray-core as bridge.


### Run
```bash
git clone https://github.com/MrMohebi/xray-iran-bridge-configs.git
cd xray-iran-bridge-configs
```
then run:
```bash
docker compose up -d --build
```

### Connecting configs

socks5 :
> proxy all --> `socks5://myuser:123654789@SERVER_IP:9901`
> 
> not ir --> `socks5://myuser:123654789@SERVER_IP:9902`
> 
> proxy all --> `socks5://myuser:123654789@SERVER_IP:9900`

vless:
> not ir --> `vless://1ca7eeef-fd8b-4eef-8127-8583e44943bc@server_ip:7700?security=&type=tcp&encryption=none#my-proxy_not-ir`


### configs:
you can change the default passwords and usernames from [master inbounds file](configs-master/inbounds.json)

### Routing rules
There are some pre-config routing rules

- `bypass iranian websites`: It won't proxy iranian websites and just pass them from bridge iran server
    ###### **Note:** Base on *.ir domains and provided list from [iran-hosted-domains](https://github.com/bootmortis/iran-hosted-domains)
- `just bans`: It will only proxy censored and sanctioned websites
    ###### **Note:** Base on [domain-list-iran-bans](https://github.com/MrMohebi/domain-list-iran-bans)
- `proxy all`: I think no explanation is required.



### Inbounds
- socks5
    - not ir
    - just bans
    - proxy all
- vless:
    - not ir
