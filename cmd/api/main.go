package main

import (
	"log"

	"github.com/Gorpetrosov/golang-social/internal/db"
	"github.com/Gorpetrosov/golang-social/internal/env"
	"github.com/Gorpetrosov/golang-social/internal/store"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}

	cfg := config{
		addr: env.GetEnvString("ADDR", ":8000"),
		db: dbConfig{
			addr: env.GetEnvString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOpenConns: env.GetEnvInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleConns: env.GetEnvInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime: env.GetEnvString("DB_MAX_IDLE_TIME", "30m"),
		},
	}

    db, err := db.New(cfg.db.addr, cfg.db.maxOpenConns, cfg.db.maxOpenConns, cfg.db.maxIdleTime)


	if(err != nil) {
		log.Panic(err)
	}

	defer db.Close();

	log.Println("database connection pool established")

	store := store.NewPostgresStorage(db)

	app := &application{
		config: cfg,
		store: store,
	}


	mux := app.mount()

	log.Fatal(app.run(mux))
}
