package main

import (
	"log"

	"github.com/BaldurDevs/baldur_go-library/pkg/goutils/configsloader"
	"github.com/BaldurDevs/baldur_go-library/pkg/http/baserest"
	"github.com/gin-gonic/gin"
)

func main() {

	loader := configsloader.ConfigsLoaderFactory()
	err := loader.LoadFile(".env")
	if err != nil {
		return
	}

	if err := run(); err != nil {
		log.Fatalf("error running application: %v", err)
	}
}

func run() error {
	router := gin.New()

	pingHandler := baserest.NewPingHandler()
	pingHandler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		return err
	}

	return nil
}
