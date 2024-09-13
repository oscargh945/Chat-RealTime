package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/oscargh945/go-Chat/infrastructure/http/handler"
)

var r *gin.Engine

func RouterInit(userHandler *handler.UserHandler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUserHandler)
	r.POST("/login", userHandler.LoginHandler)
	r.POST("/logout", userHandler.LogoutHandler)
	r.POST("/refresh-token", userHandler.RefreshTokensHandler)
}

func Init(addr string) error {
	return r.Run(addr)
}
