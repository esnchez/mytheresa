package main

import "log"

func main() {

	cfg := config{
		addr: ":8080",
	}

	app := &app{
		config: cfg,
	}

	mux := app.NewMux()

	log.Fatal(app.start(mux))
}