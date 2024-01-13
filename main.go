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

	//? Manejador para servir los archivos estáticos
	fs := http.FileServer(http.Dir("static"))

	//? Ruta para acceder a los archivos estáticos
	router.Handle("/static/", http.StripPrefix("/static/", fs))

	//? rutas y handler - función encargada de manejar una solicitud HTTP
	router.HandleFunc("/", handlers.Index)
	router.HandleFunc("/new", handlers.NewGame)
	router.HandleFunc("/start", handlers.Game)
	router.HandleFunc("/play", handlers.Play)
	router.HandleFunc("/about", handlers.About)

	port := ":4000"
	log.Printf("Servidor escuchando en http://localhost%s\n", port)

	//? Iniciar el servidor, nil para el manejador de rutas
	log.Fatal(http.ListenAndServe(port, router))
}
