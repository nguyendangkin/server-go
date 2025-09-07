package model

import "time"

// mặc định nó sẽ thêm s, chẳng hạn như User, khi thật là users. Và chúng viết thường cho tất cả
type User struct {
	ID        uint      `gorm:"primaryKey"`      // khóa chính tự tăng
	Username  string    `gorm:"unique;not null"` // tên đăng nhập, duy nhất
	Email     string    `gorm:"unique;not null"` // email, duy nhất
	Password  string    `gorm:"not null"`        // hash, không lưu raw text
	CreatedAt time.Time // gorm tự quản lý
	UpdatedAt time.Time // gorm tự quản lý
}
