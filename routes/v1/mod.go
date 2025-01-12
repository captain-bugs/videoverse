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
	group.GET("/share/view/", response.GinWrapper(pc.GetViewFile, ph.GetViewFile))

	user := group.Group("/user/").Use(middleware.Auth())
	{
		user.GET("/", response.GinWrapper(pc.GetUser, ph.GetUser))
	}

	video := group.Group("/video/").Use(middleware.Auth())
	{
		video.GET("/:id/", response.GinWrapper(pc.GetVideo, ph.GetVideo))
		video.POST("/", response.GinWrapper(pc.PostVideo, ph.PostVideo))
		video.POST("/trim/", response.GinWrapper(pc.PostTrimVideo, ph.PostTrimVideo))
		video.POST("/merge/", response.GinWrapper(pc.PostMergeVideo, ph.PostMergeVideo))
	}

	share := group.Group("/share/").Use(middleware.Auth())
	{
		share.GET("/video/:id/", response.GinWrapper(pc.GetGenerateShareLink, ph.GetShare))
	}

}
