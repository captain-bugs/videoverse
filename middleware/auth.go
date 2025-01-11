package middleware

import (
	"github.com/gin-gonic/gin"
	"videoverse/pkg/auth"
	"videoverse/pkg/logbox"
	"videoverse/response"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if len(token) == 0 {
			token = c.GetHeader("Authorisation")
		}

		if len(token) == 0 {
			reply := response.UnAuthorized("token length 0")
			c.JSON(reply.StatusCode, reply)
			c.Abort()
			return
		}

		manager := auth.NewTokenManager()
		details, err := manager.VerifyToken(&token)
		if err != nil {
			logbox.NewLogBox().Debug().Err(err).Str("token", token).Msg("failed to verify token")
			reply := response.UnAuthorized(err.Error())
			c.JSON(reply.StatusCode, reply)
			c.Abort()
			return
		}
		c.Set("user_id", details.UserID)
		c.Next()
	}
}
