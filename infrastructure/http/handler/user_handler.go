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

func (h *UserHandler) LoginHandler(c *gin.Context) {
	var user entities.LoginRequest
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u, err := h.UserService.Login(c.Request.Context(), &user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.SetCookie("jwt", u.AccessToken, 3600, "/", "localhost", false, true)

	response := &entities.LoginResponse{
		AccessToken:  u.AccessToken,
		RefreshToken: u.RefreshToken,
	}

	c.JSON(http.StatusOK, response)
}

func (h *UserHandler) LogoutHandler(c *gin.Context) {
	c.SetCookie("jwt", "", -1, "", "", false, true)
	c.JSON(http.StatusOK, gin.H{"Messages": "Logout successfully!"})
}

func (h *UserHandler) RefreshTokensHandler(c *gin.Context) {
	var request struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tokens, err := h.UserService.RefreshTokenUserService(request.RefreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Invalid refresh token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  tokens.AccessToken,
		"refresh_token": tokens.RefreshToken,
	})
}
