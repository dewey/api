# Mullvad

https://play.google.com/store/apps/details?id=net.mullvad.mullvadvpn

must use API 26 or higher.

- https://github.com/httptoolkit/frida-android-unpinning/issues/35
- https://github.com/mullvad/mullvadvpn-app/issues/4862
- https://github.com/mullvad/mullvadvpn-app/issues/4863

## type 14

~~~
{
   "active": 8,
   "city_code": 5,
   "city_name": 6,
   "country_code": 3,
   "country_name": 4,
   "fqdn": 7,
   "hostname": 2,
   "ipv4_addr_in": 11,
   "ipv6_addr_in": 12,
   "network_port_speed": 13,
   "owned": 9,
   "provider": 10,
   "status_messages": 15,
   "stboot": 8,
   "type": 14
},
"al-tia-ovpn-001",
"al",
"Albania",
"tia",
"Tirana",
"al-tia-ovpn-001.relays.mullvad.net",
true,
false,
"iRegister",
"31.171.154.50",
"2a04:27c0:0:4::a01f",
1,
"openvpn",
[],
~~~

## type 25

~~~
{
   "hostname": 17,
   "country_code": 3,
   "country_name": 4,
   "city_code": 5,
   "city_name": 6,
   "fqdn": 18,
   "active": 8,
   "owned": 9,
   "provider": 10,
   "ipv4_addr_in": 19,
   "ipv6_addr_in": 20,
   "network_port_speed": 13,
   "stboot": 8,
   "pubkey": 21,
   "multihop_port": 22,
   "socks_name": 23,
   "socks_port": 24,
   "type": 25,
   "status_messages": 26
},
"al-tia-wg-001",
"al-tia-wg-001.relays.mullvad.net",
"31.171.153.66",
"2a04:27c0:0:3::a01f",
"bPfJDdgBXlY4w3ACs68zOMMhLUbbzktCKnLOFHqbxl4=",
3155,
"al-tia-wg-socks5-001.relays.mullvad.net",
1080,
"wireguard",
[],
~~~

## type 256

~~~
{
   "hostname": 245,
   "country_code": 246,
   "country_name": 247,
   "city_code": 248,
   "city_name": 249,
   "fqdn": 250,
   "active": 8,
   "owned": 9,
   "provider": 164,
   "ipv4_addr_in": 251,
   "ipv6_addr_in": 252,
   "network_port_speed": 13,
   "stboot": 8,
   "ipv4_v2ray": 253,
   "ssh_fingerprint_sha256": 254,
   "ssh_fingerprint_md5": 255,
   "type": 256,
   "status_messages": 257
},
"be-bru-br-001",
"be",
"Belgium",
"bru",
"Brussels",
"be-bru-br-001.relays.mullvad.net",
"37.120.218.170",
null,
"37.120.218.171",
"SHA256:V4I8U519hTW7c/KkyTpz3sSGBb8X9B14UoFyga2DKPk",
"MD5:cf:ed:2d:b4:70:9d:9b:3b:c4:d2:36:26:c7:a6:06:1f",
"bridge",
[],
~~~
