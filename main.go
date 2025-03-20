package main

import "fmt"
type Player struct {
	name   string
	health int
	stamina int
	isDefending bool
}

const (
	damage = 10
	staminaCost = 5
	specialDamage = 1.5
	defense = .5
)

func attack(attacker, defender *Player) {
	if defender.isDefending {
		defender.health -= damage * defense
		defender.isDefending = false
	} else {
		defender.health -= damage
	}
	println("Attacker:", attacker.name, "Defender:", defender.name, "Health:", defender.health)
}

func defend(player *Player) {
	player.isDefending = true
}

// special move ignores defense
func specialMove(attacker,defender *Player) {
	if attacker.stamina - staminaCost > 0 {
		defender.health -= damage * specialDamage
		attacker.stamina /= staminaCost
		fmt.Println("Attacker:", attacker.name, "Defender:", defender.name, "Health:", defender.health)
	} else{
		fmt.Println("Not enough stamina")
	}
}


func getChoice() int {
	var choice int
	_, err :=fmt.Scanln(&choice)
	if err != nil{
		println("Invalid input")
		return -1
	}
	return choice
}



func main() {
	player1 := Player{"Player 1", 100, 100, false}
	player2 := Player{"Player 2", 100, 100, false}

	for player1.health > 0 && player2.health > 0 {
		var choice int

		println("Player 1:", player1.name, "Health:", player1.health, "Stamina:", player1.stamina, "Defending:", player1.isDefending)
		println("Player 2:", player2.name, "Health:", player2.health, "Stamina:", player2.stamina, "Defending:", player2.isDefending)

		println("Player 1's turn, select 1,2 or 3 for your next move...\n1. Attack\n2. Defend\n3. Special Move")
		choice = getChoice()

		switch choice {
			case 1:
				fmt.Println("Attack")
				attack(&player1, &player2)
			case 2:
				fmt.Println("Defend")
				defend(&player1)
			case 3:
				fmt.Println("Special Move")
				specialMove(&player1, &player2)
			case -1:
				continue
		}

		println("Player 1:", player1.name, "Health:", player1.health, "Stamina:", player1.stamina, "Defending:", player1.isDefending)
		println("Player 2:", player2.name, "Health:", player2.health, "Stamina:", player2.stamina, "Defending:", player2.isDefending)
		
		if(player2.health <= 0) {
			println("Player 1 wins!")
			break
		}

		println("Player 2's turn, select 1,2 or 3 for your next move...")
		choice = getChoice()

		switch choice {
			case 1:
				attack(&player2, &player1)
			case 2:
				defend(&player2)
			case 3:
				specialMove(&player2, &player1)
			case -1:
				continue
		}

		println("Player 1:", player1.name, "Health:", player1.health, "Stamina:", player1.stamina, "Defending:", player1.isDefending)
		println("Player 2:", player2.name, "Health:", player2.health, "Stamina:", player2.stamina, "Defending:", player2.isDefending)

		if(player1.health <= 0) {
			println("Player 2 wins!")
			break
		}

	}
}