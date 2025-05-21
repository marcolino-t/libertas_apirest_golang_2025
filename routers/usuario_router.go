package routers

import (
	"apigolang/controllers"

	"github.com/gorilla/mux"
)

func SetupRouterUsuario(router *mux.Router) {
	router.HandleFunc("/usuarios", controllers.GetUsuarios).Methods("GET")
	router.HandleFunc("/usuarios/{id}", controllers.GetUsuarioById).Methods("GET")
	router.HandleFunc("/usuarios", controllers.CreateUsuario).Methods("POST")
	router.HandleFunc("/usuarios/{id}", controllers.UpdateUsuario).Methods("PUT")
	router.HandleFunc("/usuarios/{id}", controllers.DeleteUsuario).Methods("DELETE")
}
