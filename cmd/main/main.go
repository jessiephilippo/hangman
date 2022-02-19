package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var (
	usedLetters    = []string{}
	guessedLetters = map[rune]bool{}
	dictionary     = []string{
		"Hello world",
		"Zimbabwe",
		"Counter strike",
		"Game project",
		"Horse",
		"School",
		"hangman",
		"Shooter",
		"Terrorism",
		"National Geographic",
		"Snake",
		"Optimus Prime",
		"Computer",
		"Motor Cycle",
	}
)

func main() {
	rand.Seed(time.Now().UnixNano())
	targetWord := strings.ToLower(dictionary[rand.Intn(len(dictionary))])

	hangmanState := 0

	fmt.Printf("Word has %d letters\n", len(targetWord))

	for !isGameOver(targetWord, guessedLetters, hangmanState) {
		printGameState(targetWord, guessedLetters, hangmanState)

		input := readInput()
		if len(input) != 1 {
			fmt.Println("Invalid input. please use letters only")
			continue
		}

		fmt.Printf("\nLetters used: %v\n\n", usedLetters)

		letter := rune(input[0])
		if strings.ContainsRune(targetWord, letter) {
			guessedLetters[letter] = true
		} else {
			hangmanState++
		}
	}

	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("Invalid state. Game is over and there is no winner!")
	}
}

func printGameState(targetWord string, guessedLetters map[rune]bool, hangmanState int) {
	fmt.Println(getWordGuessingProgress(targetWord, guessedLetters))
	fmt.Println()
	fmt.Println(printHangman(hangmanState))
}

func getWordGuessingProgress(targetWord string, guessedLetters map[rune]bool) string {
	result := ""

	for _, v := range targetWord {
		if v == ' ' {
			result += " "
		} else if guessedLetters[v] {
			result += fmt.Sprintf("%c", v)
		} else {
			result += "_"
		}
	}

	return result
}

func printHangman(hangmanState int) string {
	data, err := ioutil.ReadFile(fmt.Sprintf("states/hangman%d", hangmanState))
	if err != nil {
		panic(err)
	}

	return string(data)
}

func readInput() string {
	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	usedLetters = append(usedLetters, input)

	return strings.TrimSpace(input)
}

func isWordGuessed(targetWord string, guessedLetters map[rune]bool) bool {
	for _, v := range targetWord {
		if !guessedLetters[v] {
			return false
		}
	}

	return true
}

func isHangmanComplete(hangmanState int) bool {
	return hangmanState >= 9
}

func isGameOver(targetWord string, guessedLetters map[rune]bool, hangmanState int) bool {
	return isWordGuessed(targetWord, guessedLetters) || isHangmanComplete(hangmanState)
}
