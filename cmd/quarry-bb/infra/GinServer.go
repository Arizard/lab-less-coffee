package infra

import (
	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/core"
	"github.com/gin-gonic/gin"
)

type GinServer struct{}

func NewGinServer() *GinServer {
	return &GinServer{}
}

func (srv *GinServer) Start(core *core.Core) {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1.GET("/ping", func(c *gin.Context) {
		c.JSON(
			200,
			gin.H{
				"message": "pong",
			},
		)
	})

	v1.POST("/user/register", func(c *gin.Context) {
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
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		user.PasswordHash = ""

		c.JSON(200, user)
	})

	r.Run(":8090")
}
