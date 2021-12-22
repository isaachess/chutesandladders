package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	err := run()
	if err != nil {
		log.Fatal(err)
	}
}

func bobby() {

}

func run() error {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Choose your goal (1-6): ")
	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	goal, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return err
	}

	if goal > 6 {
		return fmt.Errorf("goal must be between 1 and 6")
	}

	fmt.Print("How many players are there? ")

	text, err = reader.ReadString('\n')
	if err != nil {
		return err
	}

	numPlayers, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return err
	}

	players, err := readPlayerNames(numPlayers, reader)
	if err != nil {
		return err
	}

	fmt.Println("The names of the players are:", players)
	err = startGame(players, goal)
	if err != nil {
		return err
	}

	return nil
}

func startGame(players []string, goal int) error {
	for {
		for _, player := range players {
			won := takeTurn(player, goal)
			if won {
				fmt.Printf("Player %s won!\n", player)
				return nil
			}
		}
	}
}

func takeTurn(player string, goal int) (won bool) {
	fmt.Printf("It is %s's turn!\n", player)
	roll := rollD6()
	fmt.Printf("%s rolled a %d\n", player, roll)
	if roll == goal {
		return true
	}
	return false
}

func rollD6() int {
	return rollDie(6)
}

func rollDie(sides int) int {
	return rand.Intn(sides) + 1
}

func readPlayerNames(numPlayers int, reader *bufio.Reader) ([]string, error) {
	var names []string
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Please enter the name of player %d: ", i+1)

		text, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		names = append(names, strings.TrimSpace(text))
	}
	return names, nil
}
