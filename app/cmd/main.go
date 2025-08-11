package main

import (
	"demo/internal/app"
	"demo/internal/database"
	"demo/internal/handlers"
	"demo/internal/kafka"
	"demo/internal/repository"
	"demo/internal/utils"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	db := database.NewDB("postgres://postgres:postgres@db:5432/advertisements?sslmode=disable")
	orderRepository := repository.NewOrderRepository(db.DB_CONN)
	orderHandler := handlers.NewOrderHandler(orderRepository)
	app := app.NewApp(db, orderHandler)

	kafkaCfg := kafka.KafkaInfo{
		BrokkerAddress: "kafka:9092",
		Topic:          "orders",
		GroupId:        "users",
	}
	os.Stdout.Sync()

	go func() {
		app.MustStart()
	}()

	go func(kafkaCfg *kafka.KafkaInfo) {
		kafka.NewKafka(kafkaCfg)
	}(&kafkaCfg)

	go func() {
		time.Sleep(10 * time.Second)
		utils.StartSender()
	}()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
	<-stop

	app.Stop()
	log.Println("Gracefully stopped")
}

//TODO
// Реализовать бэкенд
// 2) получение конкретного объявления GET /order/id
// 3) получение всех объявлений GET /orders
// 4) Реализовать кэширование данных в сервисе
// 5) Написать + связать фронтенд и бэк
// 6) Тесты?
// 7) Упаковать readme
