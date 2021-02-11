[![Run on Repl.it](https://repl.it/badge/github/azaelgg/meteoric)](https://repl.it/github/azaelgg/meteoric) [![views](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/azaelgg/meteoric)](https://hits.seeyoufarm.com)

<p align="center">
    <img src="meteoric.gif" alt="animated"/>
</p>

*still wip, so if any errors open issue and specify.*

## Installation
  - [git](https://git-scm.com/download)
  - [golang](https://golang.org/dl/)
  - run these commands:
```
go get github.com/valyala/fasthttp 
go get github.com/Jeffail/gabs 
go get github.com/gorilla/websocket 
go get github.com/gookit/color 
go get gopkg.in/yaml.v2 
go get gopkg.in/mgo.v2/bson
``` or head to `install.bat` then `run.bat`

## Features 
  - webhook logging
  - fast claim time
  - its own gateway connection to speed things up

## Setting up
```yaml
# meteoric-sniper config.yaml file

generalConfig: {
  token: 'main_token',
  bot: false,
  logging: {
    webhook_url: 'webhook_url',
    file_name: '/data/meteoric-logs.txt'
  },

  envConfig: {
    use_env_file: true,
    token_value_name: 'TOKEN'
  }
}
```
   - fill your [config.yaml file](https://github.com/azaelgg/meteoric/blob/main/config.yaml)

## Todo
  - Snipe from bot account to user account.
  - ~Webhook logging.~
  - ~Terminal ascii & color.~
  - ~Make connection to gateway more stable.~
  - Multiple alt account support.
  - ~Linux support (?)~
  - [href](https://github.com/azaelgg/meteoric/blob/main/discord/gateway.go#L57)
