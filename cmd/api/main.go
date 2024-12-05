package main

import (
	"log"

	"github.com/Gorpetrosov/golang-social/internal/env"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
	  log.Fatal("Error loading .env file")
	}
  

	cfg := config{
		addr: env.GetEnvString("ADDR", ":8000"),
	}
	app := &application{
		config: cfg,
	}

	mux := app.mount()

	log.Fatal(app.run(mux))
}
