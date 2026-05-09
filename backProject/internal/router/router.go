package router

import (
	"github.com/gin-gonic/gin"
	"zyqHome/backProject/internal/handler"
	"zyqHome/backProject/internal/middleware"
)

func Setup(
	authHandler *handler.AuthHandler,
	menuHandler *handler.MenuHandler,
	iconHandler *handler.IconHandler,
	configHandler *handler.ConfigHandler,
	behaviorHandler *handler.BehaviorHandler,
	uploadHandler *handler.UploadHandler,
	imageHandler *handler.ImageHandler,
	jwtSecret string,
) *gin.Engine {
	r := gin.Default()

	r.Use(middleware.CORSMiddleware())

	// Public API (no auth required)
	public := r.Group("/api")
	{
		// Site config
		public.GET("/config", configHandler.GetPublicConfig)

		// Icons
		public.GET("/icons", iconHandler.GetActiveIcons)

		// Menus
		public.GET("/menus", menuHandler.GetActiveMenus)

		// Behavior tracking
		public.POST("/behavior/track", behaviorHandler.Track)

		// Auth
		public.POST("/auth/login", authHandler.Login)
	}

	// Static files
	r.Static("/uploads", "./uploads")

	// Admin API (auth required)
	admin := r.Group("/api/admin")
	admin.Use(middleware.AuthMiddleware(jwtSecret))
	{
		// Auth
		admin.GET("/auth/me", authHandler.Me)

		// Menus
		admin.GET("/menus", menuHandler.GetAllMenus)
		admin.POST("/menus", menuHandler.Create)
		admin.PUT("/menus/:id", menuHandler.Update)
		admin.DELETE("/menus/:id", menuHandler.Delete)
		admin.PUT("/menus/sort", menuHandler.BatchSort)

		// Icons
		admin.GET("/icons", iconHandler.GetAllIcons)
		admin.POST("/icons", iconHandler.Create)
		admin.PUT("/icons/:id", iconHandler.Update)
		admin.DELETE("/icons/:id", iconHandler.Delete)
		admin.PUT("/icons/sort", iconHandler.BatchSort)

		// Config
		admin.GET("/config", configHandler.GetAllConfigs)
		admin.PUT("/config", configHandler.UpdateConfigs)

		// Images
		admin.GET("/images", imageHandler.GetList)
		admin.GET("/images/categories", imageHandler.GetCategories)
		admin.POST("/images", imageHandler.Upload)
		admin.PUT("/images/:id", imageHandler.Update)
		admin.DELETE("/images/:id", imageHandler.Delete)

		// Upload
		admin.POST("/upload", uploadHandler.Upload)

		// Dashboard
		admin.GET("/dashboard", behaviorHandler.Dashboard)
	}

	return r
}
