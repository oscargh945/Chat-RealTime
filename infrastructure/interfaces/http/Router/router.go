package Router

import (
	"github.com/gin-gonic/gin"
	"github.com/oscargh945/go-Chat/infrastructure/interfaces/http/handler"
	"github.com/oscargh945/go-Chat/infrastructure/interfaces/http/webSocket"
)

var r *gin.Engine

func RouterInit(userHandler *handler.UserHandler, wsHandler *webSocket.Handler) {
	r = gin.Default()

	r.POST("/signup", userHandler.CreateUserHandler)
	r.POST("/login", userHandler.LoginHandler)
	r.POST("/logout", userHandler.LogoutHandler)
	r.POST("/refresh-token", userHandler.RefreshTokensHandler)

	r.POST("ws/createRoom/", wsHandler.CreateRoom)
	r.GET("/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/ws/getRooms/", wsHandler.GetRooms)
	r.GET("/ws/getClients/:roomsId", wsHandler.GetClients)
}

func Init(addr string) error {
	return r.Run(addr)
}
