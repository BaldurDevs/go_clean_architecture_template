package main

import (
	"github.com/BaldurDevs/go_clean_architecture_templatego_clean_architecture_template/internal/entrypoint/handler/rest"
	"log"

	"github.com/BaldurDevs/baldur_go-library/pkg/goutils/configsloader"
	"github.com/BaldurDevs/baldur_go-library/pkg/http/baserest"
	"github.com/gin-gonic/gin"
)

func main() {

	loader := configsloader.ConfigsEnvLoaderFactory()
	err := loader.LoadFile()
	if err != nil {
		log.Fatalf("error loading env file: %v", err)
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

	helloHandler := rest.HelloHandlerFactory()
	helloHandler.RegisterRoutes(router)

	if err := router.Run(); err != nil {
		return err
	}

	return nil
}
