package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"unicode"

	"github.com/tjarratt/babble"
)

var usedLetters = []string{}

func main() {
	targetWord := randomWord()
	hangmanState := 0

	guessedLetters := map[rune]bool{}

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
	fmt.Println("Game over... word was:", targetWord)

	if isWordGuessed(targetWord, guessedLetters) {
		fmt.Println("You win!")
	} else if isHangmanComplete(hangmanState) {
		fmt.Println("You lose!")
	} else {
		panic("Invalid state. Game is over and there is no winner!")
	}
}

func randomWord() string {
	babbler := babble.NewBabbler()
	randomWord := babbler.Babble()
	replacer := strings.NewReplacer("-", " ", "'", "")

	return replacer.Replace(randomWord)
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
		}

		if guessedLetters[unicode.ToLower(v)] {
			result += fmt.Sprintf("%c", v)
		}

		result += "_"
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
		if !guessedLetters[unicode.ToLower(v)] {
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
