package main

import (
	"log"
	"net/http"

	// importando de otra carpeta
	"rpsweb/handlers"
)

func main() {

	//? crear enrutador
	router := http.NewServeMux()

	// rutas y handler - función encargada de manejar una solicitud HTTP
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/start", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":8080"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)

	// iniciar el servidor, nil para el manejador de rutas
	log.Fatal(http.ListenAndServe(port, router))
}
