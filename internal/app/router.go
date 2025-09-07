package app

import (
	"chin_server/internal/handler"
	"chin_server/internal/repository"
	"chin_server/internal/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()

	// 1. Khởi tạo repository
	userRepo := repository.NewUserRepository(db)
	// 2. Khởi tạo service (cần repo)
	userService := service.NewUserService(userRepo)
	// 3. Khởi tạo handler (cần service)
	userHandler := handler.NewUserHandler(userService)

	// 4. Đăng ký route
	api := r.Group("/api/v1")
	{
		api.POST("/register", userHandler.Register)
	}
	return r
}
