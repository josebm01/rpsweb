package handlers

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

//? Rutas
const (
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)



//? handler - función encargada de manejar una solicitud HTTP
func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Inicio")
	renderTemplate( w, "index.html", nil )

}

func NewGame(w http.ResponseWriter, r *http.Request) {
	renderTemplate( w, "new-game.html", nil )
}

func Game(w http.ResponseWriter, r *http.Request) {
	renderTemplate( w, "game.html", nil )
}

func Play(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Jugar")
}

func About(w http.ResponseWriter, r *http.Request) {
	renderTemplate( w, "about.html", nil )
}




//! función genérica para renderizar plantillas
/**
	- recibe responseWriter
	- base de pagina - ruta base desde templates/archivo.html
	- nombre de la pagina html 
	- data - información
**/
func renderTemplate( w http.ResponseWriter, page string, data any) {
	//? Cargando plantillas HTML
	// Must - envuelve función que se encarga de cargar plantillas
	tpl := template.Must(template.ParseFiles( templateBase, templateDir + page ))

	// Renderizando plantilla y validando error
	err := tpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		http.Error(w, "Error al renderizar plantillas", http.StatusInternalServerError)
		log.Println(err)
		return
	}
}