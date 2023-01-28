package services

import (
	"github.com/timoyo93/auth-backend/pkg/auth"
	"github.com/timoyo93/auth-backend/pkg/db"
	"github.com/timoyo93/auth-backend/pkg/models"
)

type UserRepository interface {
	GetUser(username string) (*models.UserDB, error)
	GetUserByAccessToken(token string) (*models.UserDB, error)
	AddUser(user *models.UserDB) error
	CheckForAccessToken(token string) error
	SetAccessTokenForUser(username, token string) error
	RemoveAccessTokenForUser(token string) error
}

type AuthService struct {
	repo UserRepository
}

func InitAuth(repo *db.Repository) AuthService {
	return AuthService{repo: repo}
}

func (s AuthService) GetUser(username string) (*models.UserDB, error) {
	user, err := s.repo.GetUser(username)
	return user, err
}

func (s AuthService) GetUserByAccessToken(token string) (*models.UserDB, error) {
	user, err := s.repo.GetUserByAccessToken(token)
	return user, err
}

func (s AuthService) CreateUser(request *models.AuthRequest) error {
	hash, err := auth.HashPassword(request.Password)
	if err != nil {
		return err
	}
	dbUser := models.NewDBUser(request.Username, hash)
	if err = s.repo.AddUser(dbUser); err != nil {
		return err
	}
	return nil
}

func (s AuthService) CheckAccessToken(token string) error {
	if err := s.repo.CheckForAccessToken(token); err != nil {
		return err
	}
	return nil
}

func (s AuthService) SetAccessToken(username string) (string, error) {
	token := auth.TokenGenerator()
	err := s.repo.SetAccessTokenForUser(username, token)
	return token, err
}

func (s AuthService) RemoveAccessToken(token string) error {
	if err := s.repo.RemoveAccessTokenForUser(token); err != nil {
		return err
	}
	return nil
}
