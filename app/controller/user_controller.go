package app

import (
	"net/http"
	"servicev0/src/application"
	"servicev0/src/domain/entity"

	"github.com/gin-gonic/gin"
)

//Users struct defines the dependencies that will be used
type Users struct {
	us application.UserAppInterface
}

//NewUsers constructor
func NewUsers(us application.UserAppInterface) *Users {
	return &Users{
		us: us,
	}
}

//SaveUser : Save new user
func (s *Users) SaveUser(c *gin.Context) {
	var user entity.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_json": "invalid json",
		})
		return
	}
	//validate the request:
	validateErr := user.Validate("")
	if len(validateErr) > 0 {
		c.JSON(http.StatusUnprocessableEntity, validateErr)
		return
	}
	newUser, err := s.us.SaveUser(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusCreated, newUser.PublicUser())
}

//GetUsers : Get Data ALl User
func (s *Users) GetUsers(c *gin.Context) {
	var filter entity.User
	if err := c.ShouldBindQuery(&filter); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"invalid_filter": "invalid param filter",
		})
		return
	}

	//Get user profile by user Login
	user, err := s.us.GetUsers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, user)
}
