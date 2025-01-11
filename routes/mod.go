package routes

import (
	"net/http"
	"time"
	"videoverse/pkg/config"
	"videoverse/pkg/logbox"
	"videoverse/repository"
	"videoverse/routes/v1"

	"github.com/gin-gonic/gin"
)

type Router struct {
	router *gin.Engine
}

func health() gin.HandlerFunc {
	return func(c *gin.Context) {
		health := map[string]interface{}{
			"version":   config.BACKEND_VERSION,
			"service":   config.SERVICE_NAME,
			"message":   "ok",
			"timestamp": time.Now().UTC().UnixMilli(),
		}
		c.JSON(200, health)
		return
	}
}

func notFound() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusNotFound, map[string]interface{}{
			"code":      "PAGE_NOT_FOUND",
			"message":   "page not found",
			"timestamp": time.Now().UTC().UnixMilli(),
		})
		return
	}
}

func cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE, PATCH")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func NewRouter() *Router {
	gin.DefaultWriter = logbox.NewLogBox()
	return &Router{gin.New()}
}

func (r *Router) configure() {
	r.enableCORS()
	r.enableRecover()
	r.routerHealth()
	r.router.NoRoute(notFound())
}

func (r *Router) routerHealth() {
	r.router.GET("/health/", health())
}

func (r *Router) enableCORS() {
	r.router.Use(cors())
}

func (r *Router) enableRecover() {
	r.router.Use(gin.Recovery())
}

func (r *Router) setV1Routes(repo repository.IRepo) {
	v1.Register(r.router.Group("api/v1/"), repo)
}

func (r *Router) SetRoutes(repo repository.IRepo) http.Handler {
	r.configure()
	r.setV1Routes(repo)
	return r.router
}
