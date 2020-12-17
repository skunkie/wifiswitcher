# OpenWrt WiFi Switcher

## Getting started

1. Run in the project's root folder:

```sh
$ make prepare
$ make build
```
The compiled binary can be found in the `bin` directory.

2. Create a file [config.yml](./config_example.yml):

```
host: 192.168.1.1
port: 22
username: root
password: openwrt
```
