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
	return readUntil('\n')
}

func readUntil(delimeter byte) (string, error) {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString(delimeter)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(text), nil
}

func startGame(players []string, goal int) error {
	for {
		for _, player := range players {
			won, err := takeTurn(player, goal)
			if err != nil {
				return err
			}
			if won {
				fmt.Printf("Player %s won!\n", player)
				return nil
			}
		}
	}
}

func takeTurn(player string, goal int) (won bool, err error) {
	fmt.Printf("\n\nIt is %s's turn!\n\nPress enter/return to roll!\n\n", player)
	_, err = readLine()
	if err != nil {
		return false, err
	}
	for i := 0; i < 3; i++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Print(". ")
	}
	time.Sleep(500 * time.Millisecond)
	fmt.Print("\n")
	roll := rollD6()
	fmt.Printf("%s rolled a %d\n", player, roll)
	return roll == goal, nil
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
