package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/middlewares"
	"github.com/Paulo-Ariel-Pareja/backend-red-social-go/routes"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

/*Manejadores handler para rutas, cors y server port*/
func Manejadores() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(routes.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(routes.Login)).Methods("POST")
	router.HandleFunc("/verperfil", middlewares.ChequeoBD(middlewares.ValidarJWT(routes.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificar-perfil", middlewares.ChequeoBD(middlewares.ValidarJWT(routes.ModificarPerfil))).Methods("PUT")
	router.HandleFunc("/guardar-tweet", middlewares.ChequeoBD(middlewares.ValidarJWT(routes.InsertarTweet))).Methods("POST")
	router.HandleFunc("/leet-tweet", middlewares.ChequeoBD(middlewares.ValidarJWT(routes.LeoTweets))).Methods("GET")
	router.HandleFunc("/eliminar-tweet", middlewares.ChequeoBD(middlewares.ValidarJWT(routes.EliminarTweet))).Methods("DELETE")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
