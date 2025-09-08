package app

import (
	"chin_server/internal/handler"
	"chin_server/internal/middleware"
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

	// 4. Middleware
	authMiddleware := middleware.AuthMiddleware(userService)

	// Public routes (không cần token)
	public := r.Group("/api/v1")
	{
		public.POST("/register", userHandler.Register)
		public.POST("/login", authMiddleware.LoginHandler)
		public.GET("/refresh_token", authMiddleware.RefreshHandler)
	}

	// Protected routes (cần token)
	protected := r.Group("/api/v1")
	protected.Use(authMiddleware.MiddlewareFunc())
	{
		protected.GET("/me", userHandler.Me)
	}

	return r
}
