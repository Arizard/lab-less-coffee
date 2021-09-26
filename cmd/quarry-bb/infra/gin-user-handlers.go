package infra

import (
	"errors"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/core"
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/entity"
	"github.com/gin-gonic/gin"
)

func pingHandlerGin(core *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "pong",
			},
		)
	}
}

func protectedPingHandlerGin(core *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "protected pong",
			},
		)
	}
}

func userRegisterHandlerGin(core *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {

		var body struct {
			Login    string
			Password string
		}

		err := c.ShouldBindJSON(&body)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		if body.Login == "" || body.Password == "" {
			c.JSON(400, gin.H{
				"error": "login and password must not be empty",
			})
			return
		}

		user, err := core.RegisterUser(body.Login, body.Password)

		if err != nil {
			c.JSON(translateError(err), gin.H{"error": err.Error()})
			return
		}

		if user == nil {
			c.JSON(500, gin.H{"error": "user was nil"})
			return
		}

		user.PasswordHash = ""

		c.JSON(200, user)
	}
}

func getUserHandlerGin(core *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		params := c.Request.URL.Query()

		login, username, uid := params.Get("login"), params.Get("username"), params.Get("uid")

		if login == "" {
			login = username
		}

		var user *entity.User
		var coreErr error

		if login != "" {
			user, coreErr = core.GetUserByLogin(login)
		}
		if uid != "" {
			user, coreErr = core.GetUserByUID(entity.UserUID(uid))
		}

		if login == "" && uid == "" {
			c.AbortWithStatus(400)
			return
		}

		if coreErr != nil {
			c.AbortWithStatusJSON(translateError(coreErr), gin.H{"error": coreErr.Error()})
			return
		}

		if user == nil {
			c.AbortWithStatus(500)
		}

		user.PasswordHash = ""

		c.JSON(200, user)

		return

	}
}

var statusByDomainError = map[error]int{
	entity.ErrorNotFound: 404,
}

func translateError(domainError error) int {
	for err, code := range statusByDomainError {
		if errors.Is(domainError, err) {
			return code
		}
	}

	return 500
}
