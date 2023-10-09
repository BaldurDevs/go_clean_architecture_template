package rest

import (
	"github.com/BaldurDevs/baldur_go-library/pkg/goutils/mapper"
	"github.com/BaldurDevs/baldur_go-library/pkg/http/baserest"
	"github.com/BaldurDevs/go_clean_architecture_templatego_clean_architecture_template/internal/entrypoint/handler/contracts"
	"github.com/gin-gonic/gin"
)

func HelloHandlerFactory() baserest.Handler {
	return NewHelloHandler()
}

func NewHelloHandler() baserest.Handler {
	return &helloHandler{}
}

type helloHandler struct {
	autoMapper *mapper.Mapper
}

func (handler *helloHandler) RegisterRoutes(c *gin.Engine) {
	c.GET(basePath+"hello", handler.helloOperation)
}

func (handler *helloHandler) helloOperation(c *gin.Context) {

}

func (handler *helloHandler) getDataFromRequest(c *gin.Context) (*contracts.HelloRequest, error) {
	helloRequestContract := &contracts.HelloRequest{}
	if err := c.BindJSON(helloRequestContract); err != nil {
		return nil, err
	}

	return helloRequestContract, nil
}
