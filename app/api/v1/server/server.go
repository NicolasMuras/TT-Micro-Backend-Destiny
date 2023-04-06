package server

import (
	"TT-Micro-Backend-Destiny/api/v1/handlers"
	routers "TT-Micro-Backend-Destiny/api/v1/router"
	"TT-Micro-Backend-Destiny/pkg/cache"
	"TT-Micro-Backend-Destiny/pkg/db"
	"TT-Micro-Backend-Destiny/pkg/repository"
	"TT-Micro-Backend-Destiny/pkg/service"

	"log"

	"go.mongodb.org/mongo-driver/mongo"
)

// Server define the server dependencies & attributes.
type Server struct {
	router *routers.Router
	db     *mongo.Collection
}

// StartServer starts the API server with all the dependencies running.
func StartServer(port string) {
	s := Server{
		db: db.GetCollection(),
	}
	DestinyRepository := repository.NewDestinyRepository(s.db)
	//Initialize services
	DestinyService := service.NewDestinyService(DestinyRepository)
	// Initialize redis
	redisCache := cache.ConnectRedis()
	// handler creation
	handler := handlers.NewHandler(DestinyService, redisCache)

	// Server router
	s.router = routers.NewRouter(handler)
	err := s.router.Engine.Run(port)
	if err != nil {
		log.Fatal("Error: Server connection closed")
	}
}
