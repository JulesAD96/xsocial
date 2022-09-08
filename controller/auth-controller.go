package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//AuthCOntroller interface is a contract what this controller can do
type AuthCOntroller interface {
	Login(ctx *gin.Context)
	Register(ctx *gin.Context)
}

type authCOntroller struct {
	//this is where you put your service
}

// NewAuthController creates a new instance of AuthControll
func NewAuthController() AuthCOntroller {
	return &authCOntroller{}
}

func (c *authCOntroller) Login(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Login hello"})
}

func (c *authCOntroller) Register(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"message": "Register hello"})
}
