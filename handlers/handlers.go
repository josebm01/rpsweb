package handlers

import (
	"fmt"
	"net/http"
)

// handler - función encargada de manejar una solicitud HTTP
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Pagina de Inicio")
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Crear nuevo juego")
}

func Game(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Juego")
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Acerca de")
}



