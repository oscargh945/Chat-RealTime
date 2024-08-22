package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/oscargh945/go-Chat/domain/entities"
	"github.com/oscargh945/go-Chat/domain/service"
	"net/http"
)

type UserHandler struct {
	UserService service.UserService
}

func NewUserHandler(s service.UserService) *UserHandler {
	return &UserHandler{
		UserService: s,
	}
}

func (h *UserHandler) CreateUserHandler(c *gin.Context) {
	var user entities.CreateUserReq
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	res, err := h.UserService.CreateUserService(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, res)
}
