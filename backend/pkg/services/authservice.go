package services

import (
	"github.com/timoyo93/auth-backend/pkg/db"
	"github.com/timoyo93/auth-backend/pkg/models"
	"github.com/timoyo93/auth-backend/pkg/utils"
)

type UserRepository interface {
	GetUser(username string) (*models.UserDb, error)
	GetUserByAccessToken(token string) (*models.UserDb, error)
	AddUser(user *models.UserDb) error
	CheckForAccessToken(token string) bool
	SetAccessTokenForUser(username, token string) error
	RemoveAccessTokenForUser(token string) (bool, error)
}

type AuthService struct {
	repo UserRepository
}

func InitAuth(repo *db.Repository) AuthService {
	return AuthService{repo: repo}
}

func (s AuthService) GetUser(username string) (*models.UserDb, error) {
	user, err := s.repo.GetUser(username)
	return user, err
}

func (s AuthService) GetUserByAccessToken(token string) (*models.UserDb, error) {
	user, err := s.repo.GetUserByAccessToken(token)
	return user, err
}

func (s AuthService) CreateUser(auth *models.AuthRequest) error {
	hash, err := utils.HashPassword(auth.Password)
	if err != nil {
		return err
	}
	dbUser := models.NewDbUser(auth.Username, hash)
	if err = s.repo.AddUser(dbUser); err != nil {
		return err
	}
	return nil
}

func (s AuthService) CheckAccessToken(token string) bool {
	success := s.repo.CheckForAccessToken(token)
	return success
}

func (s AuthService) SetAccessToken(username string) (string, error) {
	token := utils.TokenGenerator()
	err := s.repo.SetAccessTokenForUser(username, token)
	return token, err
}

func (s AuthService) RemoveAccessToken(token string) (bool, error) {
	ok, err := s.repo.RemoveAccessTokenForUser(token)
	return ok, err
}
