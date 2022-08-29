package authhandler

import (
	"github.com/labstack/echo/v4"
	"github.com/timoyo93/auth-backend/pkg/db"
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
func RegisterUser(c echo.Context) error {
	u := new(models.AuthRequest)
	if err := c.Bind(u); err != nil {
		return err
	}
	hash, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	dbu := models.UserDb{
		Username: u.Username,
		Hash:     hash,
	}
	a, _ := db.GetUser(u.Username)
	if a != nil {
		return c.JSON(http.StatusBadRequest, "User already existing")
	}
	if err := db.AddUser(dbu); err != nil {
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
func LoginUser(c echo.Context) error {
	u := new(models.AuthRequest)
	if err := c.Bind(u); err != nil {
		return err
	}
	user, err := db.GetUser(u.Username)
	if err != nil {
		if strings.Contains(err.Error(), "no documents in result") {
			return c.JSON(http.StatusNotFound, "No user found")
		}
		return c.JSON(http.StatusBadRequest, err)
	}
	if user == nil {
		return c.JSON(http.StatusNotFound, "No user found")
	}
	ok := utils.CheckPasswordHash(u.Password, user.Hash)
	if !ok {
		return c.JSON(http.StatusUnauthorized, "Username or Password not matching")
	}
	user.Token = utils.TokenGenerator()
	err, successful := db.SetAccessTokenForUser(*user)
	if err != nil || !successful {
		return c.JSON(http.StatusBadRequest, err)
	}
	cookie := &http.Cookie{
		Name:  "JSESSIONID",
		Value: user.Token,
		Path:  "/",
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "Successfully logged in")
}

// LogOutUser godoc
// @Summary Log out
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} string "Successfully logged out"
// @Failure 404 {string} string
// @Failure 400 {string} string
// @Router /auth/logout [post]
func LogOutUser(c echo.Context) error {
	cookie, err := c.Cookie("JSESSIONID")
	if err != nil {
		return err
	}
	if cookie == nil {
		return c.JSON(http.StatusBadRequest, "No cookie found")
	}
	err, ok := db.RemoveAccessTokenForUser(cookie.Value)
	if !ok {
		return c.JSON(http.StatusBadRequest, "")
	}
	cookie = &http.Cookie{
		Name:   "JSESSIONID",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, "Removed cookie, you are signed out now")
}

// TriggerAuth godoc
// @Summary Check if authentication is still active
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {string} string
// @Failure 401 {string} string
// @Router /auth [get]
func TriggerAuth(c echo.Context) error {
	return nil
}
