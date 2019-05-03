# Galileo

## Описание

Сервис регистрации показаний температуры

## Конфигурация

Конфигурация сервиса через переменные окружения.<br>
Доступные параметры конфигурации
```
// версия приложения
GALILEO_APP_VERSION
// адрес приложения
GALILEO_LOCAL_HOST 
// порт приложения
GALILEO_LOCAL_PORT
// признак отображения timestamp в логах
GALILEO_SHOW_LOG_TIME
// путь для дампа хранилища
GALILEO_DUMP_PATH 
```

## Установка

Клонировать репозиторий проекта:

    $ go get https://github.com/panace9i/galileo.git
Перейти в директорию проекта:

    $ cd $GOPATH/src/github.com/panace9i/galileo/
Выполнить команду:
    
    $ make build
Установить переменные окружения:
  
    $ export GALILEO_APP_VERSION=1.0.0
    $ export GALILEO_LOCAL_HOST=127.0.0.1
    $ export GALILEO_LOCAL_PORT=8089 
    $ export GALILEO_SHOW_LOG_TIME=true
    $ export GALILEO_DUMP_PATH=/tmp/galileo_dumps
Запустить тесты:

    $ make test

## Запуск
Выполнить команду:
    
    $ galileo

## Доступные методы
Регистрация пользователя

    POST: /users/registration
Регистрация устройства  

    POST: /devices/registration
    
Cтатистика по устройству
    
    POST: /devices
    
Пользовательские устройства    

    GET: /users/devices
    
Cтатистика по устройству    

    GET: /devices/:id
    
Информация о сервисе 

    GET: /info

## Примеры запросов
    $ curl -X POST -H "Content-Type: application/json" -d '{"email":"asd2"}' 127.0.0.1:8089/users/registration
    $ curl -X POST -H "Content-Type: application/json" -d '{"email":"asd2","name":"test"}' 127.0.0.1:8089/devices/registration
    $ curl -X POST -H "Content-Type: application/json" -H "Authorization: 7dc822e25adcf9a9a9ec287423e787b2" -d '[{"time":"2019-02-02 00:00:00","temp":100}]' 127.0.0.1:8089/devices
    $ curl -X GET -H "Authorization: a67995ad3ec084cb38d32725fd73d9a3" 127.0.0.1:8089/users/devices
    $ curl -X GET -H "Authorization: a67995ad3ec084cb38d32725fd73d9a3" 127.0.0.1:8089/devices/7dc822e25adcf9a9a9ec287423e787b2



