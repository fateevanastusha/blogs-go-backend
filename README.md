# blogs-go-backend

REST API для блогов и постов на Go.

## Требования

- Go 1.26+

## Установка

```sh
go mod download
cp .env.example .env
```

## Запуск

```sh
make dev     # hot reload (air)
make build   # сборка бинаря ./server
./server
make test    # тесты
```

## Переменные окружения

| Переменная      | Описание                          | По умолчанию |
| --------------- | --------------------------------- | ------------ |
| `AUTH_LOGIN`    | логин для basic auth              | —            |
| `AUTH_PASSWORD` | пароль для basic auth             | —            |
| `PORT`          | порт сервера                      | `3005`       |
| `LOG_LEVEL`     | `debug` / `info` / `warn` / `error` | `info`     |

