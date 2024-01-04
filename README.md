# Iran bridge Xray config generator
will generate `outbounds.json` base on provided proxies from [`proxies_active_no_403_under_1000ms.txt`](https://raw.githubusercontent.com/MrMohebi/xray-proxy-grabber-telegram/master/collected-proxies/xray-json/actives_no_403_under_1000ms.txt) 
and update hosted xray-core as bridge.


### Run
first run: `docker compose up -d --build`

then: `docker exec -id xray-core /root/xray-core/xray-iran-bridge-updater`

It will update it's self every 30 minutes. Change it [here](./docker-compose.yml#L26). 

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
