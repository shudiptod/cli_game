package main

import "fmt"
type Player struct {
	name   string
	health int
}

func handleAttack(attacker, defender *Player) {
	defender.health -= 10
	println("Attacker:", attacker.name, "Defender:", defender.name, "Health:", defender.health)
}

func main() {
	player1 := Player{"Player 1", 100}
	player2 := Player{"Player 2", 100}

	for player1.health > 0 && player2.health > 0 {
		println("Player 1's turn, press enter to attack...")
		fmt.Scanln()
		handleAttack(&player1, &player2)
		
		if(player2.health <= 0) {
			println("Player 1 wins!")
			break
		}

		println("Player 2's turn, press enter to attack...")
		fmt.Scanln()
		handleAttack(&player2, &player1)
		if(player1.health <= 0) {
			println("Player 2 wins!")
			break
		}
		
	}
}