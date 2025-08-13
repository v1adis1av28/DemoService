package main

import (
	"demo/internal/app"
	"demo/internal/cache"
	"demo/internal/config"
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
	cfg := config.GetConfig("config/dev.yml")

	db := database.NewDB(cfg.Database.PostgresURL)
	redisClient := cache.NewRedisClient(cfg.Redis.Address, cfg.Redis.Password)

	orderRepository := repository.NewOrderRepository(db.DB_CONN)
	orderService := service.NewOrderService(orderRepository, redisClient)
	orderHandler := handlers.NewOrderHandler(orderService)

	app := app.NewApp(db, orderHandler, cfg)

	go func() {
		kafka.NewKafka(&kafka.KafkaInfo{
			BrokkerAddress: cfg.Kafka.BrokerAddress,
			Topic:          cfg.Kafka.Topic,
			GroupId:        cfg.Kafka.GroupId,
		})
	}()

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
// Добавить чтение из файла конфигурации
// 6) Тесты
// 7) Упаковать readme
