package api

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/models"
	"github.com/timoyo93/auth-backend/pkg/utils"
	"net/http"
	"strings"
)

// RegisterUser godoc
// @Summary Creates a new User
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.AuthRequest true "Register User"
// @Success 200 {string} string "user created"
// @Router /auth/register [post]
func (a API) RegisterUser(c echo.Context) error {
	var auth *models.AuthRequest
	if err := c.Bind(&auth); err != nil {
		return err
	}
	user, _ := a.authService.GetUser(auth.Username)
	if user != nil {
		return c.JSON(http.StatusBadRequest, "User already existing")
	}
	if err := a.authService.CreateUser(auth); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, "User was created successfully")
}

// LoginUser godoc
// @Summary Log in with given user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param user body models.AuthRequest true "Login User"
// @Success 200 {object} string "Logged in"
// @Failure 404 {string} string
// @Failure 400 {string} string
// @Router /auth/login [post]
func (a API) LoginUser(c echo.Context) error {
	var auth *models.AuthRequest
	if err := c.Bind(&auth); err != nil {
		return err
	}
	user, err := a.authService.GetUser(auth.Username)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return c.JSON(http.StatusNotFound, "No user found")
		}
		return c.JSON(http.StatusBadRequest, err)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "No user found")
	}
	if ok := utils.CheckPasswordHash(auth.Password, user.Hash); !ok {
		return c.JSON(http.StatusUnauthorized, "Username or Password not matching")
	}
	token, err := a.authService.SetAccessToken(auth.Username)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookie := &http.Cookie{
		Name:  CookieName,
		Value: token,
		Path:  CookiePath,
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "Successfully logged in")
}

// LogoutUser godoc
// @Summary Log out
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} string "Successfully logged out"
// @Failure 404 {string} string
// @Failure 400 {string} string
// @Router /auth/logout [post]
func (a API) LogoutUser(c echo.Context) error {
	cookie, err := c.Cookie(CookieName)
	if err != nil {
		return err
	}
	if cookie == nil {
		return c.JSON(http.StatusBadRequest, "No cookie found")
	}
	ok, err := a.authService.RemoveAccessToken(cookie.Value)
	if !ok {
		return c.JSON(http.StatusBadRequest, "Could not remove token")
	}
	cookie.Value = ""
	cookie.MaxAge = -1
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "Successfully logged out")
}

// TriggerAuth godoc
// @Summary Check if authentication is still active
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /auth [get]
func (a API) TriggerAuth(_ echo.Context) error {
	return nil
}
