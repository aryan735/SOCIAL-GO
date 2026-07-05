package main

import (
	"log"

	"github.com/aryan735/SOCIAL-GO/internal/env"
	"github.com/aryan735/SOCIAL-GO/internal/store"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")

	}
	cfg := config{addr: env.GetString("ADDR", ":8080")}
	store:= store.NewStorage(nil)
	app := application{config: cfg, store: store}


	mux := app.mount()
	log.Fatal(app.run(mux))
}
