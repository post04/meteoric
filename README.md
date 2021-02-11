[![Run on Repl.it](https://repl.it/badge/github/azaelgg/meteoric)](https://repl.it/github/azaelgg/meteoric) [![views](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/azaelgg/meteoric)](https://hits.seeyoufarm.com)

<p align="center">
    <img src="meteoric.gif" alt="animated"/>
</p>

## Installation
  - [git](https://git-scm.com/download)
  - [golang](https://golang.org/dl/)
  - Once done, open your terminal and run this command: `go get github.com/valyala/fasthttp && go get github.com/Jeffail/gabs && go get github.com/gorilla/websocket && go get github.com/gookit/color && go get gopkg.in/yaml.v2`, and wait for it to finish.

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

## Todo
  - Snipe from bot account to user account.
  - ~Webhook logging.~
  - ~Terminal ascii & color.~
  - ~Make connection to gateway more stable.~
  - Multiple alt account support.
  - ~Linux support (?)~
