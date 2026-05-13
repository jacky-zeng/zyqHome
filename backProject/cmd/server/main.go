package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"

	"zyqHome/backProject/internal/config"
	"zyqHome/backProject/internal/database"
	"zyqHome/backProject/internal/handler"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
	"zyqHome/backProject/internal/router"
	"zyqHome/backProject/internal/service"
)

func main() {
	// Load .env file if exists
	_ = godotenv.Load()

	cfg := config.Load()

	// Init database connections
	database.InitMySQL(cfg)
	database.InitRedis(cfg)

	// Seed default admin user if not exists
	seedDefaultAdmin()

	// Migrate existing users: set empty role to "admin" (backward compatibility)
	database.DB.Model(&model.User{}).Where("role = '' OR role IS NULL").Update("role", "admin")

	// Init repositories
	userRepo := repository.NewUserRepo()
	menuRepo := repository.NewMenuRepo()
	iconRepo := repository.NewIconRepo()
	configRepo := repository.NewConfigRepo()
	behaviorRepo := repository.NewBehaviorRepo()
	imageRepo := repository.NewImageRepo()
	userIconRepo := repository.NewUserIconRepo()

	// Init services
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	menuService := service.NewMenuService(menuRepo)
	iconService := service.NewIconService(iconRepo)
	configService := service.NewConfigService(configRepo)
	behaviorService := service.NewBehaviorService(behaviorRepo)
	imageService := service.NewImageService(imageRepo, "./uploads")
	userIconService := service.NewUserIconService(userIconRepo)

	// Init handlers
	authHandler := handler.NewAuthHandler(authService)
	menuHandler := handler.NewMenuHandler(menuService)
	iconHandler := handler.NewIconHandler(iconService)
	configHandler := handler.NewConfigHandler(configService)
	behaviorHandler := handler.NewBehaviorHandler(behaviorService)
	uploadHandler := handler.NewUploadHandler("./uploads")
	imageHandler := handler.NewImageHandler(imageService, "./uploads")
	userIconHandler := handler.NewUserIconHandler(userIconService)
	userImageHandler := handler.NewUserImageHandler(imageService, "./uploads")

	// Setup router
	r := router.Setup(
		authHandler,
		menuHandler,
		iconHandler,
		configHandler,
		behaviorHandler,
		uploadHandler,
		imageHandler,
		userIconHandler,
		userImageHandler,
		cfg.JWTSecret,
	)

	addr := fmt.Sprintf(":%s", cfg.ServerPort)
	log.Printf("Server starting on %s", addr)
	if err := r.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func seedDefaultAdmin() {
	var count int64
	database.DB.Model(&model.User{}).Count(&count)
	if count == 0 {
		hashedPassword, err := service.HashPassword("admin123")
		if err != nil {
			log.Fatalf("Failed to hash password: %v", err)
		}
		admin := &model.User{
			Username: "admin",
			Password: hashedPassword,
			Nickname: "管理员",
			Role:     "admin",
		}
		if err := database.DB.Create(admin).Error; err != nil {
			log.Fatalf("Failed to seed admin: %v", err)
		}
		log.Println("Default admin user created (admin/admin123)")
	}
}
