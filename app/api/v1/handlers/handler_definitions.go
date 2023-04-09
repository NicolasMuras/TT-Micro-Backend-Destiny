package handlers

import (
	"TT-Micro-Backend-Destiny/pkg/cache"
	"TT-Micro-Backend-Destiny/pkg/service"

	"github.com/gin-gonic/gin"
)

// Handler contains attributes for the object handlers.
type Handler struct {
	DestinyService service.DestinyInterface
	redisCache     cache.RedisCache
}

// DestinyHandler define the methods for the Destiny object.
type DestinyHandler interface {
	Retrieve(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
	Create(ctx *gin.Context)
	List(ctx *gin.Context)
	RetrieveGroup(ctx *gin.Context)
}

// Interface define the methods for the handler.
type Interface interface {
	DestinyHandler
}

// NewHandler Return the Handlder with all the handler dependencies.
func NewHandler(DestinyService service.DestinyInterface, redisCache cache.RedisCache) Interface {
	return &Handler{
		DestinyService: DestinyService,
		redisCache:     redisCache,
	}
}
