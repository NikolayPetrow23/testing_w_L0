# Сервис поиска заказов

Добро пожаловать в репозиторий сервиса поиска заказов WBTECH TEST 0! Этот проект является тестовым заданием стажировки, 
созданный с использованием Go, PostreSQL, HTML, CSS, Bootstrap. Здесь вы найдете описание проекта, инструкции по 
запуску и развертыванию, а также другую полезную информацию.

## Описание проекта

WBTECH TEST 0 предоставляет возможность поиска заказов в базе данных (кэшэ).


## Основные функции
- **Поиск заказа:** заказ можно найти по OrderUid.
- **Добавление заказа:** Заказ можно добавить передав сообщение в Nats Streaming в виде json'a.

## Технологии

- **Go - 1.21**
- **PostgreSQL - 13.3**
- **Docker - 4.19**

## Установка и запуск

```bash
# Установка репозитория
git clone git@github.com:NikolayPetrow23/testing_w_L0.git
```

```bash
# Поднятие базы данных с помощью Docker
docker run -p 5432:5432 --name "Имя вашей БД" -e POSTGRES_USER="Введите пользователя для БД" -e POSTGRES_PASSWORD="Введите пароль для БД" -e POSTGRES_DB="Имя вашей БД" -d postgres:13.3
```

```bash
# Поднятие Nats Streaming в Docker
docker run -d --name "Имя вашего контейнера Nats" -p 4222:4222 -p 6222:6222 -p 8222:8222
```

```bash
# Запуск проекта
go run cmd/main.go
```

## Автор
[Николай Петров](https://github.com/NikolayPetrow23)