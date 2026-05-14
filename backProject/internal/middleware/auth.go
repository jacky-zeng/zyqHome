package middleware

import (
	"strings"

	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/model"
	jwtpkg "zyqHome/backProject/pkg/jwt"
	"zyqHome/backProject/pkg/response"
)

func AuthMiddleware(jwtSecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Unauthorized(c, "未提供认证令牌")
			c.Abort()
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Unauthorized(c, "认证令牌格式错误")
			c.Abort()
			return
		}

		claims, err := jwtpkg.ParseToken(parts[1], jwtSecret)
		if err != nil {
			response.Unauthorized(c, "认证令牌无效或已过期")
			c.Abort()
			return
		}

		// 检查用户是否被禁用
		var user model.User
		if err := database.DB.First(&user, claims.UserID).Error; err != nil || user.Status != 1 {
			response.Unauthorized(c, "账号已被禁用")
			c.Abort()
			return
		}

		c.Set("user_id", claims.UserID)
		c.Set("username", claims.Username)
		c.Set("role", claims.Role)
		c.Next()
	}
}

func AdminMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role, _ := c.Get("role")
		if role != "admin" {
			response.Forbidden(c, "无权限访问")
			c.Abort()
			return
		}
		c.Next()
	}
}
