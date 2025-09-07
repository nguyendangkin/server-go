package app

import (
	"chin_server/config"
	"fmt"
)

func Run() {
	// 1. khởi chạy load config
	cfg := config.LoadConfig()

	// 2. Dùng config đó để kết nối database
	db := config.ConnectDatabase(cfg) // trả về db

	// khởi tạo gin
	// truyền db vào router
	r := SetupRouter(db)
	// lấy port
	port := fmt.Sprintf(":%d", cfg.Server.Port)

	// khởi chạy với port tùy chỉnh
	r.Run(port)
}
