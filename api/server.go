package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/scipiia/effectivemobiletask/db/sqlc"
)

type Server struct {
	store  db.Store
	router *gin.Engine
	//kafka   worker.Process
}

func NewServer(store db.Store) (*Server, error) {

	// p, err := worker.NewKafkaProduce(context.Background(), "localhost:9092", "FIO")
	// if err != nil {
	// 	return nil, fmt.Errorf("cannot start produce in servir %w", err)
	// }
	server := &Server{
		store: store,
		//kafka:   p, //todo
	}
	router := gin.Default()

	router.POST("/users", server.createUser)
	router.GET("/users/:id", server.getUser)
	router.GET("/users", server.ListUser)
	router.DELETE("/users/:id", server.deleteUser)
	router.PATCH("/users", server.updateUser)

	server.router = router
	return server, nil
}

func (server *Server) Start(addres string) error {
	return server.router.Run(addres)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
