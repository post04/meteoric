[![Run on Repl.it](https://repl.it/badge/github/azaelgg/meteoric)](https://repl.it/github/azaelgg/meteoric) [![views](https://hits.seeyoufarm.com/api/count/incr/badge.svg?url=https://github.com/azaelgg/meteoric)](https://hits.seeyoufarm.com)

<p align="center">
    <img src="meteoric.gif" alt="animated"/>
</p>

## Installation
  - [git](https://git-scm.com/download)
  - [golang](https://golang.org/dl/)
  - Once done, open your terminal and run this command: `go get github.com/sadlil/go-trigger && go get github.com/valyala/fasthttp && go get github.com/Jeffail/gabs && go get github.com/gorilla/websocket`, and wait for it to finish.

## Setting up
    ```yaml
config:
  token: "main_token"
  secondary_token: "alt_token"
  bot: false
  snipe_to_main_token: false

  ignore_codes: 
    - "xxx"
  ignore_guilds:
    - 123
  ignore_users:
    - 123

  webhook_url: ""
    ```

## Todo
  - Snipe from bot account to user account.
  - Webhook logging.
  - Terminal ascii & color.
  - Make connection to gateway more stable.
  - Multiple alt account support.
  - Linux support (?)
