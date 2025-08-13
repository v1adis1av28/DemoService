package main

import (
	"demo/internal/app"
	"demo/internal/cache"
	"demo/internal/database"
	"demo/internal/handlers"
	"demo/internal/kafka"
	"demo/internal/repository"
	"demo/internal/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	db := database.NewDB("postgres://postgres:postgres@db:5432/advertisements?sslmode=disable")

	RedisClient := cache.NewRedisClient("redis:6379", "")

	orderRepository := repository.NewOrderRepository(db.DB_CONN)
	orderService := service.NewOrderService(orderRepository, RedisClient)
	orderHandler := handlers.NewOrderHandler(orderService)

	app := app.NewApp(db, orderHandler)

	kafkaCfg := kafka.KafkaInfo{
		BrokkerAddress: "kafka:9092",
		Topic:          "orders",
		GroupId:        "users",
	}
	os.Stdout.Sync()

	go func(kafkaCfg *kafka.KafkaInfo) {
		kafka.NewKafka(kafkaCfg)
	}(&kafkaCfg)

	go func() {
		app.MustStart()
	}()
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	app.Stop()
	log.Println("Gracefully stopped")
}

//TODO
// Вынести сендер в отдельный сервис
// Добавить чтение из файла конфигурации
// 6) Тесты
// 7) Упаковать readme
