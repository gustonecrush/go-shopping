package app

import "github.com/gustonecrush/go-shop/app/controllers"

func (server *Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}