// internal/service/auth_service.go

package service

import (
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"zcloud-bg/internal/model"
	"zcloud-bg/internal/repository"
)

// Claims 用于 JWT 认证的自定义声明结构体
type Claims struct {
	Username string `json:"username"`
	UserID   string `json:"userid"`
	jwt.StandardClaims
}

type AuthService struct {
	UserRepo *repository.UserRepository
}

// AuthenticateUser 检查提供的用户名和密码是否匹配
func (service *AuthService) AuthenticateUser(username, password string) (*model.User, error) {
	user, err := service.UserRepo.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	// 验证密码
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return user, nil
}
