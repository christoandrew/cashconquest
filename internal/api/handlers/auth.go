package handlers

import (
	"net/http"

	"github.com/christo-andrew/maybe-go/internal/api/requests"
	"github.com/christo-andrew/maybe-go/internal/api/responses"
	"github.com/christo-andrew/maybe-go/internal/models"
	"github.com/christo-andrew/maybe-go/pkg/errors"
	utils "github.com/christo-andrew/maybe-go/pkg/utils"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// LoginHandler Login godoc
// @Summary Login a user
// @Description Login a user
// @Accept json
// @Produce json
// @Param user body requests.LoginRequest true "Login Request"
// @Success 200 {object} responses.LoginResponse
// @Router /users/login [post]
// @Tags auth
func LoginHandler(c *gin.Context, db *gorm.DB) {
	var loginRequest requests.LoginRequest
	err := c.BindJSON(&loginRequest)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse{Message: err.Error()})
		return
	}
	var user models.User
	user, err = getUserByCredentials(loginRequest.Email, loginRequest.Password, db)
	if err != nil {
		c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Message: err.Error()})
		return
	}
	token := user.GenerateToken()
	response := responses.LoginResponse{
		Token: token,
		Name:  user.GetFullName(),
		Email: user.Email,
		ID:    user.ID,
	}

	c.JSON(http.StatusOK, response)
}

// The most basic authentication in the world
func getUserByCredentials(email string, password string, db *gorm.DB) (models.User, error) {
	var user models.User
	db.Where("email = ?", email).First(&user)

	if user == (models.User{}) {
		return user, errors.UserNotFoundError()
	}

	if !utils.ComparePassword(user.Password, password) {
		return user, errors.InvalidCredentialsError()
	}

	return user, nil
}
