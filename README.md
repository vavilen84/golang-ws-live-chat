# Golang WebSocket based live chat

Application includes:
- Docker (application container)
- Modules
- gorilla/websocket package


## Install Docker 

https://docs.docker.com/install/

## Install docker-compose 

https://docs.docker.com/compose/install/

## Install docker-hostmanager

https://github.com/iamluc/docker-hostmanager

Run manager

```
$ docker run -d --name docker-hostmanager --restart=always -v /var/run/docker.sock:/var/run/docker.sock -v /etc/hosts:/hosts iamluc/docker-hostmanager
```

## Env

create new.env file from .env.dist and set correct project path

## Start with Docker

```
$ cd /project/path
$ docker compose up -d --build
```

## Application URL
 
http://baseapp.golang-ws-live-chat_local/

