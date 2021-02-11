[![Views](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/azaelgg/meteoric)](https://hits.seeyoufarm.com)
[![Run on Repl.it](https://repl.it/badge/github/azaelgg/meteoric)](https://repl.it/github/azaelgg/meteoric) 
[![Run on Heroku](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/azaelgg/meteoric/tree/heroku)

<p align="center">
    <img src="meteoric.gif" alt="animated"/>
</p>

## Showcase
  <details>
    <summary>Sniper showcase</summary>

  ![](https://i.imgur.com/Lg2TNQh.gif)
  </details>
 - Credits: riptide

## Installation

- [git](https://git-scm.com/download)
- [golang](https://golang.org/dl/)
- Run `install.bat` then `run.bat` or run these commands:

```
go get github.com/valyala/fasthttp
go get github.com/Jeffail/gabs
go get github.com/gorilla/websocket
go get github.com/gookit/color
go get gopkg.in/yaml.v2
go get gopkg.in/mgo.v2/bson
```

- `go run .`

## Features

- Webhook logging.
- Fast claim time.
- Its own gateway connection to speed things up.
- Anti dupes.

## Setting up

```yaml
# meteoric-sniper config.yaml file

generalConfig: {
  token: 'main_token',         # main token.
  bot: false,                  # if the main token is a bot or not.
  logging: {
    webhook_url: '',                     # webhook to log info to, keep it empty to not do any webhook logging.
    file_name: '/data/meteoric-logs.txt' # file logging, coming soon.
  },

  envConfig: {
    use_env_file: false,       # use .env file, if you're using a public host i highly recommend it.
    token_value_name: 'TOKEN'  # key value of the token inside the .env
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

---
#### Extra
  _still wip, so if any errors open issue and specify._
