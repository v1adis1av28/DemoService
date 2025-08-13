# DemoService

Демонстрационный сервис с интерфейсом, отображающий данные о заказе.

## Структура проекта

```

.
├── app/                  # Основной бэкенд с консьюмером
│   ├── cmd/              
│   ├── config/           
│   ├── internal/         
│   ├── migrations/       
│   ├── Dockerfile
│   ├── tests/
│   └── go.mod
├── frontend/             # Фронтенд (React + Vite)
│   ├── public/
│   ├── src/
│   ├── nginx.conf
│   └── Dockerfile
├── producer/             # Сервис producer отправляющий в кафку сообщения
├── docker-compose.yml    # Основная конфигурация
├── docker-compose.test.yml    # Конфигурация приложения для тестирования
├── README.md
└── go.mod

````

---

##  Запуск проекта

### Скопируйте репозиторий
```bash
git clone https://github.com/v1adis1av28/demoservice.git
cd demoservice
````

### 2. Запустить в Docker

```bash
docker-compose up --build
```

После запуска будут доступны:

* **Backend API**: [http://localhost:8080](http://localhost:8080)
* **Frontend**: [http://localhost:3000](http://localhost:3000)
* **PostgreSQL**: `localhost:5432` (user: `postgres`, pass: `postgres`)
* **Kafka**: `localhost:9092`
* **Redis**: `localhost:6379`

---

##  Интеграционные тесты

```bash
docker compose -f docker-compose.test.yml up -d 
```

Запуск тестов:

```bash
go test ./tests -v
```
