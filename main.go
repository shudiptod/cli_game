package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
type Player struct {
	name   string
	health float64
	stamina float64
	isDefending bool
}

const (
	damage = 10.0
	staminaCost = 3.5
	specialDamage = 1.5
	defense = .5
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func attack(attacker, defender *Player) {
	damageDealt := damage
	if defender.isDefending {
		damageDealt = damage * defense
		defender.isDefending = false
	} 
	
	defender.health -= damageDealt
	
	fmt.Println(Red, attacker.name, "attacks!", Reset)
	fmt.Println(Yellow, defender.name, "takes", damageDealt, "damage.", Reset)
}

// Defend function
func defend(player *Player) {
	player.isDefending = true
	fmt.Println(Blue, player.name, "is defending and will take reduced damage next turn!", Reset)
}

// special move ignores defense
func specialMove(attacker,defender *Player) {
	newStamina := attacker.stamina - attacker.stamina/staminaCost
	if newStamina > 0  {
		damageDealt := damage * specialDamage
		defender.health -= damageDealt
		attacker.stamina = newStamina
		fmt.Println(Green, attacker.name, "uses Special Move! Deals", damageDealt, "damage!", Reset)
	} else{
		fmt.Println(Red, attacker.name, "does not have enough stamina for a Special Move.", Reset)
	}
}


func getChoice() int {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Enter your choice: ")
		input, _:= reader.ReadString('\n')
		input = strings.TrimSpace(input)

		choice, err := strconv.Atoi(input)
		if err == nil && (choice>=1 && choice<=3) {
			return choice
		}
		fmt.Println("Invalid input! Please enter 1, 2, or 3.")
	}
}


// Function to display player stats
func displayStats(player1, player2 *Player) {
	fmt.Println("\n===================================")
	fmt.Println("         📊 Player Stats 📊       ")
	fmt.Println("===================================")
	fmt.Printf("%-10s | Health: %-5.1f | Stamina: %-5.1f | Defending: %-5t\n", player1.name, player1.health, player1.stamina, player1.isDefending)
	fmt.Printf("%-10s | Health: %-5.1f | Stamina: %-5.1f | Defending: %-5t\n", player2.name, player2.health, player2.stamina, player2.isDefending)
	fmt.Println("===================================")
}


func playerTurn(player, opponent *Player) {
	fmt.Printf("\n%s's turn. Choose an action:\n1. Attack\n2. Defend\n3. Special Move\n", player.name)
	choice := getChoice()
	switch choice {
		case 1:
			attack(player, opponent)
		case 2:
			defend(player)
		case 3:
			specialMove(player, opponent)
	}
}


func main() {
	player1 := Player{"Player 1", 100, 100, false}
	player2 := Player{"Player 2", 100, 100, false}

	for player1.health > 0 && player2.health > 0 {
		displayStats(&player1, &player2)

		// player 1's turn
		playerTurn(&player1, &player2)
		
		if(player2.health <= 0) {
			fmt.Println(player1.name, "wins!")
			break
		}

		playerTurn(&player2, &player1)

		if(player1.health <= 0) {
			fmt.Println(player2.name, "wins!")
			break
		}

	}
}