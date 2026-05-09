package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"zyqHome/backProject/internal/model"
	"zyqHome/backProject/internal/repository"
	jwtpkg "zyqHome/backProject/pkg/jwt"
)

type AuthService struct {
	userRepo  *repository.UserRepo
	jwtSecret string
}

func NewAuthService(userRepo *repository.UserRepo, jwtSecret string) *AuthService {
	return &AuthService{userRepo: userRepo, jwtSecret: jwtSecret}
}

type LoginResponse struct {
	Token string      `json:"token"`
	User  *model.User `json:"user"`
}

func (s *AuthService) Login(username, password, ip string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Username, s.jwtSecret)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	_ = s.userRepo.UpdateLastLogin(user.ID, ip)

	return &LoginResponse{Token: token, User: user}, nil
}

func (s *AuthService) GetCurrentUser(userID uint) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
