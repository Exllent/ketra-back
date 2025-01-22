# Веб-сервер на Go с PostgreSQL, GORM и Gin

Этот проект представляет собой простой веб-сервер на языке Go, который использует PostgreSQL в качестве базы данных, GORM для ORM, Gin для маршрутизации и валидации данных с помощью `go-playground/validator`.

## Требования

- Go 1.16 или выше
- PostgreSQL
- Docker (опционально для запуска базы данных)

## Установка

### 1. Установка зависимостей

Перейдите в директорию проекта и установите необходимые зависимости:

```bash
go get -u github.com/gin-gonic/gin
go get -u gorm.io/gorm
go get -u gorm.io/driver/postgres
go get -u github.com/go-playground/validator/v10
