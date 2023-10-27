# Добро пожаловать в NetServDB !

Network Service with Database - это консольное приложение-сервис, предназначенное для обработки HTTP POST запросов, выполнения различных операций, и взаимодействия с Redis и PostgreSQL базами данных.

## 🚀 Доступные ручки

1. **Инкрементация значения в Redis**
    
    - Метод: POST
    - Путь: `localhost:3000/redis/incr` (если при запуске не был указан в параметрах `-host` и `-port`)
    - Обработчик: `NetServDB/controllers.RedisIncr`
    - Описание: позволяет инкрементировать значение в Redis и возвращает value
      
2. **HMAC-SHA512 Подпись**
    
    - Метод: POST
    - Путь: `localhost:3000/sign/hmacsha512`
    - Обработчик: `NetServDB/controllers.SignHMACSHA512`
    - Описание: вычисляет HMAC-SHA512 подпись и возвращает ее в виде hex строки
      
3. **Добавление пользователя в PostgreSQL**
    
    - Метод: POST
    - Путь: `localhost:3000/postgres/users`
    - Обработчик: `NetServDB/controllers.AddUser`
    - Описание: добавляет пользователя в базу данных PostgreSQL и возвращает id
    
4. **Удаление ключа в Redis** `сервисная ручка для которой потребуется basic auth`
   
    - Метод: DELETE
    - Путь: `localhost:3000/redis/del` (если при запуске не был указан в параметрах `-host` и `-port`)
    - Обработчик: `NetServDB/controllers.RedisRefresh`
    - Описание: удаляет ключ из Redis
      
6. **Обновление таблицы в PostgreSQL** `сервисная ручка для которой потребуется basic auth`
    
    - Метод: DELETE
    - Путь: `localhost:3000/postgres/users`
    - Обработчик: `NetServDB/controllers.TableRefresh`
    - Описание: обновляет таблицу в PostgreSQL

## 💡 Использованные технологии

Проект разработан с использованием следующих технологий:

### Локальная версия

- [**Gin**](https://github.com/gin-gonic/gin) - роутер
- [**logrus**](https://github.com/sirupsen/logrus) - инструмент для эффективного логирования; *все логи пишутя в logs/all.log* 
- [**testify**](https://github.com/stretchr/testify) - библиотека для обеспечения покрытия проекта тестами
- [**godotenv**](https://github.com/joho/godotenv) - библиотека для конфигурирования приложения
- [**Postgresql**](https://www.postgresql.org/) - СУБД
- [**Redis**](https://redis.io/) - БД

### Задеплоенная версия

 - [**Selectel**](https://selectel.ru/) - облачный сервер

## Перечислю доступные ручки облачной версии для удобства:
- http://94.26.237.90:3000/redis/incr `POST`
- http://94.26.237.90:3000/sign/hmacsha512 `POST`
- http://94.26.237.90:3000/postgres/users `POST`
- http://94.26.237.90:3000/redis/del `DELETE`
- http://94.26.237.90:3000/postgres/users `DELETE`
   


Все компоненты запущены в `Docker` контейнерах и объединены `Docker Compose`. Деплой руками через консоль `Selectel`.

![](https://raw.githubusercontent.com/faringet/NetServDB/master/screenshots/cloud.jpeg)


## 🛠️ Установка и использование

- Клонируйте репозиторий
`git clone` https://github.com/faringet/NetServDB.git

- Перейдите в каталог проекта
cd NetServDB

- Установите зависимости с помощью go mod
`go mod download`

- Соберите приложение
`go build`

- Для запуска выполните `./NetServDB -host <host> -port <port>` (Redis `-host` и `-port`)

### Или же можно, как пример, достучаться до api через Postman:

![](https://raw.githubusercontent.com/faringet/NetServDB/master/screenshots/postman.jpeg)
