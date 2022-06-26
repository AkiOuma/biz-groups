package main

import (
	"context"
	"flag"
	"log"

	"github.com/AkiOuma/biz-groups/internal/infrastructure/config"
	"github.com/AkiOuma/biz-groups/internal/infrastructure/constructor"
)

func main() {
	configDir := flag.String("config", "", "path config file")
	flag.Parse()
	c := config.ReadConfig(*configDir)
	app := constructor.InitApp(c)
	ctx, cancel := context.WithCancel(context.Background())
	if err := app.Serve(ctx); err != nil {
		cancel()
		log.Fatalf("server shutdown because of reason: %v", err)
	}
}
