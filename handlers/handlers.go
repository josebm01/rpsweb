package handlers

import (
	"encoding/json"
	// "fmt"
	"html/template"
	"log"
	"net/http"
	"rpsweb/rps"
	"strconv"
)

//? Rutas
const (
	templateDir = "templates/"
	templateBase = templateDir + "base.html"
)

// Estructura
type Player struct {
	Name string
}

// Instancia de la estructura
var player Player

//? handler - función encargada de manejar una solicitud HTTP
func Index(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Inicio")
	restartValue()
	renderTemplate( w, "index.html", player )
}

func NewGame(w http.ResponseWriter, r *http.Request) {
	restartValue()
	renderTemplate( w, "new-game.html", nil )
}

func Game(w http.ResponseWriter, r *http.Request) {

	//Validando el método que se envía la información desde el form
	if r.Method == "POST" {
		// valida si se ha parseado correctamente los datos
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Error parsing form", http.StatusBadRequest)
			return
		}
		
		// Obteniendo datos del name del form
		player.Name = r.Form.Get("name")
	}
	
	// Redireccionando a ruta raíz al escribir la ruta directamente /
	if player.Name == "" {
		http.Redirect(w, r, "/new", http.StatusFound)
	}

	renderTemplate( w, "game.html", player )

}

func Play(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintln(w, "Jugar")

	// obteniendo el valor de la ruta como entero
	playerChoice, _ := strconv.Atoi( r.URL.Query().Get("c"))
	// Utilizando función que devuelve la estructura
	result := rps.PlayRound( playerChoice )
	// fmt.Println( result )

	// Serializando estructura a json
	out, err := json.MarshalIndent(result, "", "    ")
	if err != nil {
		log.Println(err)
		return
	}

	// respondiendo al cliente
	// Enzabezado
	w.Header().Set("Content-Type", "application/json")
	// Resultado
	w.Write(out)
}

func About(w http.ResponseWriter, r *http.Request) {
	restartValue()
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


//! Reiniciar valores
func restartValue() {
	player.Name = ""
	rps.ComputerScore = 0
	rps.PlayerScore = 0
}