package rps

import (
	"fmt" 
	"testing"
)

func TestPayRound(t *testing.T) {
	for i := 0; i < 3; i++ {
		round := PlayRound(i)
		
		switch i {
		case 0:
			fmt.Println("El jugador eligió PIEDRA")
		case 1:
			fmt.Println("El jugador eligió PAPEL")
		case 2:
			fmt.Println("El jugador eligió TIJERA")
		}

		// imprimiendo los resultados del test
		fmt.Printf("Message: %s\n", round.Message)
		fmt.Printf("ComputerChoice: %s\n", round.ComputerChoice)
		fmt.Printf("RoundResult: %s\n", round.RoundResult)
		fmt.Printf("ComputerChoiceInt: %d\n", round.ComputerChoiceInt)
		fmt.Printf("ComputerScore: %s\n", round.ComputerScore)
		fmt.Printf("PlayerScore: %s\n", round.PlayerScore)

		fmt.Println("\n=========================")
	}
}