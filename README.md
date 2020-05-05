# avitobackend

## Как меня собрать:

Require: 
* golang v1.13+
* docker-compose

Ubuntu:
```
$ git clone  https://gitlab.com/avito-2/avitobackend -b dev
$ cd avitobackend/deploy/dev
$ docker-compose up --build
```


## О важном:
Инфа, которая, возможно неактуальна:

* Сервер запущен на 84.201.181.0:5000
* TCP-Сервер для сокетов запущен на 84.201.181.0:8100


Инфа, которая актуальна:
* На ip:port/calls/make посылаем POST
* Это инициализация звонка, в будущем будет происходить с клиента
* Чтобы инициатор получил ответ в виде sdp объекта, нужно открыть соединение с тсп (ws://ip:port/) на другом клиенте.
* Пример тестового скрипта и html лежит в папке client (там как минимум надо поменять ip). А лучше пполностью переписать.



