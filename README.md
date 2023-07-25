# Iran bridge Xray config generator
will generate `outbounds.json` base on provided proxies from [`proxies_active_no_403_under_1000ms.txt`](https://raw.githubusercontent.com/MrMohebi/xray-proxy-grabber-telegram/master/proxies_active_no_403_under_1000ms.txt) 
and update hosted xray-core as bridge.

Uses GitHub Actions to update config and hosted xray-core. 

### Routing rules
There are some pre-config routing rules

- `bypass iranian websites`: It won't proxy iranian websites and just pass them from bridge iran server
    ###### *Note: * Base on *.ir domains and provided list from [iran-hosted-domains](https://github.com/bootmortis/iran-hosted-domains)
- `just bans`: It will only proxy censored and sanctioned websites
    ###### *Note: * Base on [domain-list-iran-bans](https://github.com/MrMohebi/domain-list-iran-bans)
- `proxy all`: I think no explanation is required.



### Inbounds
- socks5
    - not ir
    - just bans
    - proxy all
- vless:
    - not ir