package v1

import (
	"github.com/gin-gonic/gin"
	v1c "videoverse/cmd/api/v1.controller"
	v1h "videoverse/cmd/api/v1.handler"
	"videoverse/middleware"
	"videoverse/pkg/logbox"
	"videoverse/repository"
	"videoverse/response"
)

func Register(group *gin.RouterGroup, repo repository.IRepo) {
	logbox.NewLogBox().Info().Str("group", "v1").Str("event", "INITIALIZING_ROUTES").Msgf("")

	pc := v1c.NewControllerV1()
	ph := v1h.NewHandlerV1(repo)

	group.POST("/user/", response.GinWrapper(pc.PostUser, ph.PostUser))

	user := group.Group("/user/").Use(middleware.Auth())
	{
		user.GET("/", response.GinWrapper(pc.GetUser, ph.GetUser))
	}

	video := group.Group("/video/").Use(middleware.Auth())
	{
		video.GET("/:id/", response.GinWrapper(pc.GetVideo, ph.GetVideo))
		video.POST("/", response.GinWrapper(pc.PostVideo, ph.PostVideo))
	}

	share := group.Group("/share/").Use(middleware.Auth())
	{
		share.POST("/", response.GinWrapper(pc.PostShare, ph.PostShare))
	}

}
