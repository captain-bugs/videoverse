package v1

import (
	"github.com/gin-gonic/gin"
	"videoverse/pkg/logbox"
)

func Register(group *gin.RouterGroup) {
	logbox.NewLogBox().Info().Str("group", "v1").Str("event", "INITIALIZING_ROUTES").Msgf("")

	grp := group.Group("/").Use()
	{
		grp.GET("/test", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "test"})
		})
	}
}
