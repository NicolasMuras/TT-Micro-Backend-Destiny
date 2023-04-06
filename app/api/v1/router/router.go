package routers

import (
	"TT-Micro-Backend-Destiny/api/v1/handlers"

	"github.com/gin-gonic/gin"
)

// Router defines all the dependencies for routing.
type Router struct {
	Engine *gin.Engine
}

// NewRouter returns a new gin Engine with all the routes.
func NewRouter(h handlers.Interface) *Router {
	router := Router{
		Engine: gin.Default(),
	}

	v1 := router.Engine.Group("")
	{
		v1.GET("/destiny/:id", func(ctx *gin.Context) {
			h.Retrieve(ctx)
		})
		v1.PUT("/destiny/:id", func(ctx *gin.Context) {
			h.Update(ctx)
		})
		v1.DELETE("/destiny/:id", func(ctx *gin.Context) {
			h.Delete(ctx)
		})
		v1.POST("/destiny", func(ctx *gin.Context) {
			h.Create(ctx)
		})
		v1.GET("/destiny", func(ctx *gin.Context) {
			h.List(ctx)
		})
	}
	return &router
}
