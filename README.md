# avitobackend

## API
https://avitocalls.docs.apiary.io/

## Building app:

Require: 
* golang v1.13+
* docker-compose

Linux/OSX:
```
$ git clone  https://gitlab.com/avito-2/avitobackend
$ cd avitobackend/deploy/dev
$ docker-compose up --build
```

#Презентация:
https://docs.google.com/presentation/d/1eLq4Zb0Kda2hTSiJtuqCZC7WdhhjMi207DOrhjJ3YWc/edit?usp=sharing

#Установка и настройка STUN сервера:
```
$ git clone https://github.com/creytiv/re.git
$ cd re
$ make
$ sudo make install

$ wget http://www.creytiv.com/pub/restund-0.4.12.tar.gz
$ tar xf restund-0.4.12.tar.gz
$ cd restund-0.4.12/
$ make
$ make install
```
 Затем отредактировать конфиг файл /etc/restund.conf
 
 Запустить сервер: 
``` 
$ restund /etc/restund.conf
```

#Установка TURN сервера:
```
$ apt-get install resiprocate-turn-server
```
Отредактировать кофиг файл /etc/reTurn/reTurnServer.config
Отредактировать список пользователей /etc/reTurn/users.txt
Запустить сервер:
```
$ service resiprocate-turn-server restart
```
