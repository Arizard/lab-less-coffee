package infra

import (
	"net/http"
	"strings"

	"github.com/arizard/lab-less-coffee/cmd/quarry-bb/core"
	"github.com/gin-gonic/gin"
)

type GinServer struct{}

func NewGinServer() *GinServer {
	return &GinServer{}
}

func getBearer(h *http.Header) string {
	authzHeader := h.Get("Authorization")

	if authzHeader == "" {
		return ""
	}

	parts := strings.Split(authzHeader, " ")

	if len(parts) != 2 {
		return ""
	}

	if parts[0] != "Bearer" {
		return ""
	}

	return parts[1]
}

func ginHasBearer(c *gin.Context) bool {
	return getBearer(&c.Request.Header) != ""
}

func sessionManagerMiddleware(core *core.Core) gin.HandlerFunc {
	return func(c *gin.Context) {
		bearer := getBearer(&c.Request.Header)

		if bearer == "" {
			c.AbortWithStatus(401)
		}

		errs := []error{}
		isValid := core.UserSessionManager.ValidateAccessToken(bearer, func(err error) { errs = append(errs, err) })

		if len(errs) != 0 || !isValid {
			errsString := []string{}
			for _, err := range errs {
				errsString = append(errsString, err.Error())
			}
			c.AbortWithStatusJSON(401, gin.H{
				"error": strings.Join(errsString, "|"),
			})
		}
	}
}

func (srv *GinServer) Start(core *core.Core) {
	r := gin.Default()

	v1 := r.Group("/api/v1")

	v1public := v1.Group("/")

	v1private := v1.Group("/")
	v1private.Use(sessionManagerMiddleware(core))

	v1public.GET("/ping", pingHandlerGin(core))
	v1private.GET("/protected-ping", protectedPingHandlerGin(core))

	v1public.POST("/register", userRegisterHandlerGin(core))

	v1private.GET("/user", getUserHandlerGin(core))

	r.Run(":8090")
}
