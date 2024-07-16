package main

import (
	"fmt"
	"os"
	"os/signal"
	"test/database"
	routes "test/handlers"

	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Запуск приложения")

	err := godotenv.Load()

	if err != nil {
		fmt.Printf("Ошибка чтения env файла. Err: %s", err)
	}

	// Создаю экземпляр БД
	db := database.NewDatabase()

	// Запускаю БД и миграцию
	connection := db.MustRun()
	db.RunMigration()

	// запускаю хендлеры
	routes.Run(connection)

	stop := make(chan os.Signal, 1)
	signal.Notify(stop)
	<-stop

	if err = connection.Close(); err != nil {
		fmt.Println("Error when closing database connection: ", err)
	}
}
