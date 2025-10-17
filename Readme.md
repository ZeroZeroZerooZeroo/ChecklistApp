# Checklist Application

[![Go Version](https://img.shields.io/badge/Go-1.25.1%2B-blue.svg)](https://golang.org)
[![PostgreSQL](https://img.shields.io/badge/PostgreSQL-Database-informational.svg)](https://www.postgresql.org)
[![Redis](https://img.shields.io/badge/Redis-Caching-DC382D.svg)](https://redis.io)
[![Kafka](https://img.shields.io/badge/Apache-Kafka-231F20.svg)](https://kafka.apache.org)
[![Docker](https://img.shields.io/badge/Docker-Containerization-2496ED.svg)](https://www.docker.com)

Микросервисное приложение для управления задачами (чек-листом) с архитектурой, построенной на трех основных сервисах: API-сервис, БД-сервис и Kafka-сервис.

---

## 📖 Оглавление

- [Checklist Application](#checklist-application)
  - [📖 Оглавление](#-оглавление)
  - [🚀 О проекте](#-о-проекте)
  - [🏗 Архитектура](#-архитектура)
  - [✨ Функциональность](#-функциональность)
  - [🛠 Технологии](#-технологии)
  - [🚀 Быстрый старт](#-быстрый-старт)
    - [Предварительные требования](#предварительные-требования)
    - [Установка и запуск](#установка-и-запуск)
  - [📚 API Документация](#-api-документация)
    - [Основные эндпоинты](#основные-эндпоинты)
    - [Примеры запросов](#примеры-запросов)
    - [Миграции базы данных](#миграции-базы-данных)

---

## 🚀 О проекте

Checklist Application — это микросервисное приложение для управления задачами, построенное на архитектуре из трех независимых сервисов. Пользователи могут создавать, просматривать, отмечать и удалять задачи, а все действия логируются через Kafka.

## 🏗 Архитектура

Приложение состоит из трех основных микросервисов:

1. **API-сервис** - принимает HTTP запросы, проксирует их в БД-сервис и отправляет события в Kafka
2. **БД-сервис** - обрабатывает gRPC запросы, работает с PostgreSQL и кеширует данные в Redis
3. **Kafka-сервис** - потребляет сообщения о действиях пользователей и логирует их в файл

## ✨ Функциональность

* **Управление задачами:**
  * Создание новых задач с заголовком и описанием
  * Просмотр списка всех задач
  * Отметка задач как выполненных
  * Удаление задач

* **Кеширование:**
  * Автоматическое кеширование списка задач в Redis
  * Настраиваемый TTL для кешированных данных
  * Инвалидация кеша при изменениях

* **Мониторинг действий:**
  * Логирование всех пользовательских действий через Kafka
  * Отслеживание времени, IP-адреса и типа действия
  * Запись логов в файл с ротацией

## 🛠 Технологии

* **Язык программирования:** [Go (Golang)](https://golang.org/)
* **База данных:** [PostgreSQL](https://www.postgresql.org/)
* **Кеширование:** [Redis](https://redis.io/)
* **Мессенджер:** [Apache Kafka](https://kafka.apache.org/)
* **Коммуникация:** [gRPC](https://grpc.io/)
* **Контейнеризация:** [Docker](https://www.docker.com/) & [Docker Compose](https://docs.docker.com/compose/)
* **Миграции БД:** [golang-migrate](https://github.com/golang-migrate/migrate)

## 🚀 Быстрый старт

### Предварительные требования

- Docker & Docker Compose

### Установка и запуск

1. **Клонирование репозитория:**
```bash
git clone https://github.com/ZeroZeroZerooZeroo/ChecklistApp.git
cd ChecklistApp
```

2. **Настройка переменных окружения:**
Создайте файл `.env` в корне проекта (пример):

```bash
# Server Configuration
SERVER_HOST=0.0.0.0
SERVER_PORT=8080

# gRPC Configuration
GRPC_HOST=databaseservice
GRPC_PORT=50051

# Database Configuration
DB_HOST=postgres
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=password
DB_NAME=checklistdb
DB_SSLMODE=disable

# Redis Configuration
REDIS_HOST=redis
REDIS_PORT=6379
REDIS_PASSWORD=
REDIS_DB=0
REDIS_TTL=300

# Kafka Configuration
KAFKA_BROKERS=kafka:9092
KAFKA_TOPIC=user-actions
KAFKA_GROUP_ID=kafka-service

# Logging
LOG_FILE_PATH=./logs/user-actions.log
```

3. **Сборка и запуск приложения:**
```bash
docker-compose up -d --build
```

Приложение будет доступно по адресу: `http://localhost:8080`

## 📚 API Документация

### Основные эндпоинты

| Метод | Путь | Описание | Тело запроса |
|-------|------|-----------|--------------|
| POST | `/create` | Создать задачу | `{"title": "string", "description": "string"}` |
| GET | `/list` | Получить список задач | - |
| DELETE | `/delete` | Удалить задачу | `{"id": number}` |
| PUT | `/done` | Отметить задачу выполненной | `{"id": number}` |

### Примеры запросов

```bash
# Создание задачи
curl -X POST http://localhost:8080/create \
  -H "Content-Type: application/json" \
  -d '{
    "title": "Изучить Go",
    "description": "Изучить основы языка программирования Go"
  }'

# Получить список задач
curl http://localhost:8080/list

# Отметить задачу выполненной
curl -X PUT http://localhost:8080/done \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'

# Удалить задачу
curl -X DELETE http://localhost:8080/delete \
  -H "Content-Type: application/json" \
  -d '{"id": 1}'
```


### Миграции базы данных

Миграции выполняются автоматически при запуске БД-сервиса через встроенный механизм golang-migrate.