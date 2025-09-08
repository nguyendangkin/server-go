package middleware

import (
	"chin_server/internal/model"
	"chin_server/internal/service"
	"log"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

var identityKey = "id"

type login struct {
	Email    string `form:"email" json:"email" binding:"required,email"`
	Password string `form:"password" json:"password" binding:"required"`
}

func AuthMiddleware(userService *service.UserService) *jwt.GinJWTMiddleware {
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("your_very_secret_key"), // ⚠️ Đổi thành key thật, lưu ở env/config
		Timeout:     time.Second * 20,
		MaxRefresh:  time.Minute * 5,
		IdentityKey: identityKey,

		// tạo payload để lưu vào token
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.ID,
					"email":     v.Email,
				}
			}
			return jwt.MapClaims{}
		},

		// trích xuất thông tin người dùng từ token trong các req
		// dùng identityKey để lấy id, sau đó truy vấn người dùng từ database để get nhiều thông tin thực
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			id, ok := claims[identityKey].(float64)
			if !ok {
				return nil
			}
			user, err := userService.GetUserByID(uint(id))
			if err != nil {
				return nil
			}
			return user
		},

		// dùng để xác thực người dùng khi đăng nhập
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals login
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}

			user, err := userService.GetUserByEmail(loginVals.Email)
			if err != nil || user == nil {
				return nil, jwt.ErrFailedAuthentication
			}

			err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginVals.Password))
			if err != nil {
				return nil, jwt.ErrFailedAuthentication
			}

			return user, nil
		},

		// kiểm tra quyền truy cập của người dùng
		Authorizator: func(data interface{}, c *gin.Context) bool {
			if _, ok := data.(*model.User); ok {
				return true
			}
			return false
		},

		// xử lý khi xác thực hoặc phân quyền thất bại
		// trả về json chưa mã lỗi và thông báo
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},

		TokenLookup:   "header: Authorization",
		TokenHeadName: "Bearer",
		TimeFunc:      time.Now,
	})

	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}
	return authMiddleware
}
