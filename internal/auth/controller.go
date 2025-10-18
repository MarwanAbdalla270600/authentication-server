package auth

import (
	"authentication-server/internal/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type controller struct {
	Service ServiceInterface
}

func NewController(service ServiceInterface) *controller {
	return &controller{
		Service: service,
	}
}

func (c *controller) Register(ctx *gin.Context) {
	var body entity.RegisterObject
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := c.Service.Register(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, response)

}

func (c *controller) Login(ctx *gin.Context) {
	var body entity.LoginObject
	if err := ctx.ShouldBindJSON(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	response, err := c.Service.Login(&body)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, response)
}
