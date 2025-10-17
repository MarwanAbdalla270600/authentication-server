package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Service *ServiceInterface
}

func NewController(service *ServiceInterface) *controller {
	return &controller{
		Service: service,
	}
}

func (c *controller) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "register")
}

func (c *controller) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, "login")
}
