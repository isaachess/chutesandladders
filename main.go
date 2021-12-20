package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	err := run()
	if err != nil {
		log.Fatal("received err", err)
	}
}

func run() error {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("How many players are there? ")

	text, err := reader.ReadString('\n')
	if err != nil {
		return err
	}

	fmt.Println("You answered: ", text)

	numPlayers, err := strconv.Atoi(strings.TrimSpace(text))
	if err != nil {
		return err
	}

	names, err := readPlayerNames(numPlayers, reader)
	fmt.Println("The names of the players are:", names)
	return nil
}

func readPlayerNames(numPlayers int, reader *bufio.Reader) ([]string, error) {
	var names []string
	for i := 0; i < numPlayers; i++ {
		fmt.Printf("Please enter the name of player %d: ", i+1)

		text, err := reader.ReadString('\n')
		if err != nil {
			return nil, err
		}
		names = append(names, text)
	}
	return names, nil
}
