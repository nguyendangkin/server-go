package handler

import (
	"chin_server/internal/dto"
	"chin_server/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// UserHandler chứa các phương thức xử lý request liên quan đến User
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler khởi tạo một UserHandler mới
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// Register xử lý yêu cầu đăng ký người dùng mới
func (h *UserHandler) Register(c *gin.Context) {
	var request dto.RegisterUserRequest

	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.userService.Register(request.Email, request.Password, request.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "user_id": user.ID})
}

func (h *UserHandler) Me(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}
