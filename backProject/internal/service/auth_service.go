package service

import (
	"errors"
	"fmt"
	"math/rand"
	"regexp"
	"strings"

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

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

func (s *AuthService) Login(username, password, ip string) (*LoginResponse, error) {
	user, err := s.userRepo.FindByUsername(username)
	if err != nil {
		// 尝试用邮箱登录
		user, err = s.userRepo.FindByEmail(username)
		if err != nil {
			return nil, errors.New("用户名或密码错误")
		}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return nil, errors.New("用户名或密码错误")
	}

	if user.Status == 0 {
		return nil, errors.New("账号已被禁用")
	}

	// 兼容旧数据：role 为空则视为 admin
	role := user.Role
	if role == "" {
		role = "admin"
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Username, role, s.jwtSecret)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	_ = s.userRepo.UpdateLastLogin(user.ID, ip)

	return &LoginResponse{Token: token, User: user}, nil
}

func (s *AuthService) Register(req RegisterRequest) (*LoginResponse, error) {
	email := strings.TrimSpace(req.Email)

	// 检查邮箱是否已注册
	exists, err := s.userRepo.ExistsByEmail(email)
	if err != nil {
		return nil, errors.New("注册失败，请稍后再试")
	}
	if exists {
		return nil, errors.New("该邮箱已被注册")
	}

	// 从邮箱前缀自动生成用户名
	username := generateUsername(email)
	exists, err = s.userRepo.ExistsByUsername(username)
	if err != nil {
		return nil, errors.New("注册失败，请稍后再试")
	}
	if exists {
		username = fmt.Sprintf("%s_%d", username, rand.Intn(9000)+1000)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("注册失败，请稍后再试")
	}

	user := &model.User{
		Username: username,
		Password: string(hashedPassword),
		Email:    email,
		Role:     "member",
		Status:   1,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, errors.New("注册失败，请稍后再试")
	}

	token, err := jwtpkg.GenerateToken(user.ID, user.Username, "member", s.jwtSecret)
	if err != nil {
		return nil, errors.New("生成令牌失败")
	}

	return &LoginResponse{Token: token, User: user}, nil
}

func generateUsername(email string) string {
	parts := strings.Split(email, "@")
	base := parts[0]
	// 只保留字母数字和下划线
	reg := regexp.MustCompile(`[^a-zA-Z0-9_]`)
	base = reg.ReplaceAllString(base, "_")
	if base == "" {
		base = "user"
	}
	return base
}

func (s *AuthService) GetCurrentUser(userID uint) (*model.User, error) {
	return s.userRepo.FindByID(userID)
}

func (s *AuthService) GetMemberList(page, pageSize int) ([]model.User, int64, error) {
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 20
	}
	return s.userRepo.FindMembers(page, pageSize)
}

func (s *AuthService) GetMemberDetail(id uint) (*model.User, error) {
	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("会员不存在")
	}
	if user.Role != "member" {
		return nil, errors.New("会员不存在")
	}
	return user, nil
}

func (s *AuthService) UpdateMemberStatus(id uint, status int) error {
	if status != 0 && status != 1 {
		return errors.New("状态值无效")
	}

	user, err := s.userRepo.FindByID(id)
	if err != nil {
		return errors.New("会员不存在")
	}
	if user.Role != "member" {
		return errors.New("只能操作会员账号")
	}

	return s.userRepo.UpdateFields(id, map[string]interface{}{
		"status": status,
	})
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
