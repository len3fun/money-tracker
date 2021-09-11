package service

import (
	"crypto/sha1"
	"fmt"
	moneytracker "github.com/len3fun/money-tracker"
	"github.com/len3fun/money-tracker/pkg/repository"
)

const salt = "rhf5498hfanu4grq8wef"

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user moneytracker.User) (int, error) {
	user.Password = generatePasswordHash(user.Password)
	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}
