package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"web-calculator/application"
)

func main() {

	// contruct application
	app := application.New()

	// create context for signal interrupt
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	// run app
	err := app.Run(ctx)

	// handle error
	if err != nil {
		fmt.Printf("error: %s\n", err)
	}
}