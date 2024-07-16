package database

import (
	"fmt"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

type Database interface {
	MustRun() *sqlx.DB
	RunMigration()
}

type Postgres struct {
	Client *sqlx.DB
}

func NewDatabase() Database {
	// Поле Client инициализирую в методе MustRun
	return &Postgres{}
}

func (p *Postgres) MustRun() *sqlx.DB {
	// Переменные окружения
	var (
		host     = os.Getenv("DB_HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("DB_USERNAME")
		password = os.Getenv("DB_PASS")
		dbname   = os.Getenv("DB_NAME")
	)

	// Создаю строку для подключения к БД и открыкаю коннект
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s",
		host, port, user, password, dbname)

	db, err := sqlx.Open("pgx", psqlInfo)

	if err != nil {
		panic(err)
	}

	// Пингую БД, если не пингуется создаю панику, т.к без работающей БД приложение тоже не будет работать
	pingErr := db.Ping()
	if pingErr != nil {
		panic(pingErr)
	}

	// Присваиваю к полю Client в структуре Postgres коннект к БД
	p.Client = db

	fmt.Println("Successfully pinnged!")

	return db
}

func (p *Postgres) RunMigration() {

	b, err := os.ReadFile("./migration.sql")

	if err != nil {
		fmt.Printf("Ошибка чтения sql файла. Err: %s", err)
	}

	_, err = p.Client.Exec(string(b))

	if err != nil {
		fmt.Printf("Ошибка чтения sql файла. Err: %s", err)
	}
}
