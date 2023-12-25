package rest

import (
	"github.com/BaldurDevs/baldur_go-library/pkg/goutils/mapper"
	"github.com/BaldurDevs/baldur_go-library/pkg/http/baserest"
	"github.com/BaldurDevs/go_clean_architecture_templatego_clean_architecture_template/internal/entrypoint/handler/contracts"
	"github.com/gin-gonic/gin"
	"log"
)

func HelloHandlerFactory() baserest.Handler {
	return newHelloHandler()
}

func newHelloHandler() baserest.Handler {
	return &helloHandler{}
}

type helloHandler struct {
	autoMapper *mapper.Mapper
}

func (handler *helloHandler) RegisterRoutes(c *gin.Engine) {
	c.POST(basePath+"hello", handler.helloOperation)
}

func (handler *helloHandler) helloOperation(_ *gin.Context) {

}

func (handler *helloHandler) getDataFromRequest(_ *gin.Context) (*contracts.HelloRequest, error) {
	log.Fatalf("%s", "Implement me...")
	return nil, nil
}
