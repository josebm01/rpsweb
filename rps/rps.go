package rps

import (
	"math/rand"
	"strconv"
)

//! lógica del juego
const (
	ROCK     = 0 // Piedra - vence a las tijeras (tijeras + 1) % 3 = 0
	PAPER    = 1 // Papel - vence a la piedra (piedra + 1) % 3 = 0
	SCISSORS = 2 // Tijeras - vence al papel (papel + 1) % 3 = 0
)

// Estructura para dar resultado de cada ronda
type Round struct {
	Message           string `json:"message"`
	ComputerChoice    string `json:"computer_choice"`
	RoundResult       string `json:"round_result"`
	ComputerChoiceInt int    `json:"computer_choice_int"`
	ComputerScore     string `json:"computer_score"`
	PlayerScore       string `json:"player_score"`
}

// slices
// Mensajes para cuando gana
var winMessages = []string{
	"¡Bien hecho!",
	"Buen trabajo",
	"Deberías comprar un boleto de lotería",
	"Ah perrooooo",
}

// Mensajes para cuando pierde
var loseMessages = []string{
	"Qué lástima",
	"Intentalo de nuevo",
	"Hoy simplemente no es tu día",
	"Ni modo bro",
}

// Mensaje de empate
var drawMessages = []string{
	"Las grandes mentes piensan igual",
	"Oh oh. Intántalo de nuevo",
	"Nadie gana, pero puedes intentarlo de nuevo",
	"¿Vas a quedarte tranquilo así?",
}

// Variables para el puntaje
var ComputerScore, PlayerScore int

func PlayRound(playerValue int) Round {
	// Obteniendo número aleatorio del 0 al 2
	computerValue := rand.Intn(3)

	var computerChoice, roundResult string 
	var computerChoiceInt int 

	switch computerValue {
	case ROCK:
		computerChoiceInt = ROCK
		computerChoice = "La computadora eligió PIEDRA"		
	case PAPER:
		computerChoiceInt = PAPER
		computerChoice = "La computadora eligió PAPEL"		
	case SCISSORS:
		computerChoiceInt = SCISSORS
		computerChoice = "La computadora eligió TIJERAS"		
	}

	// Generando número aleatorio de 0 - 2, que se usa para elegir un mensaje aleatorio
	messageInt := rand.Intn(3)

	var message string 

	if playerValue == computerValue {
		
		roundResult = "Es un EMPATE"
		// seleccionar mensaje del slice drawMessages
		message = drawMessages[ messageInt ]

	} else if playerValue == (computerValue+1)%3 {

		PlayerScore++
		roundResult = "¡El jugador gana!"
		// seleccionar mensaje de winMessages
		message = winMessages[ messageInt ]

	} else {

		ComputerScore++
		roundResult = "¡La computadora gana!"
		// seleccionar mensaje de loseMessages
		message = loseMessages[ messageInt ]

	}

	// Retornando resultado con la estructura definida
	return Round{
		Message: message,  // mensaje aleatorio
		ComputerChoice: computerChoice, // lo que eligió la COMPUTADORA
		RoundResult: roundResult, // resultado 
		ComputerChoiceInt: computerChoiceInt, // lo que eligió la COMPUTADORA en entero
		ComputerScore: strconv.Itoa(ComputerScore), // puntos de la COMPUTADORA
		PlayerScore: strconv.Itoa(PlayerScore),  // puntos del JUGADOR
	}

}