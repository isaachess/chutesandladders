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
	goal, err := askForNumber("Choose your goal (1-6): ")
	if err != nil {
		return err
	}

	if goal > 6 {
		return fmt.Errorf("goal must be between 1 and 6")
	}

	numPlayers, err := askForNumber("How many players are there? ")
	if err != nil {
		return err
	}

	players, err := readPlayerNames(numPlayers)
	if err != nil {
		return err
	}

	return startGame(players, goal)
}

func askForNumber(text string) (int, error) {
	line, err := askForLine(text)
	if err != nil {
		return 0, err
	}
	return strconv.Atoi(line)
}

func askForLine(text string) (string, error) {
	fmt.Print(text)
	return readLine()
}

func readLine() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
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

func readPlayerNames(numPlayers int) ([]string, error) {
	var names []string
	for i := 0; i < numPlayers; i++ {

		line, err := askForLine(fmt.Sprintf("Please enter the name of player %d: ", i+1))
		if err != nil {
			return nil, err
		}
		names = append(names, line)
	}
	return names, nil
}
