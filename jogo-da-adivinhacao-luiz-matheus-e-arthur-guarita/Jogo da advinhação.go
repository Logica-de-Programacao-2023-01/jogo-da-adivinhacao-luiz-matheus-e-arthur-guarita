package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxErrorsBeforeHint = 2

var language string

func main() {
	rand.Seed(time.Now().UnixNano())
	chooseLanguage()
	playGame()
}

func chooseLanguage() {
	fmt.Println("Escolha o seu idioma (Choose a language):")
	fmt.Println("1. Português (Portuguese)")
	fmt.Println("2. English (Inglês)")
	var choice int
	fmt.Scan(&choice)
	fmt.Scanln() // Consumir o caractere de nova linha

	switch choice {
	case 1:
		language = "pt"
	case 2:
		language = "en"
	default:
		fmt.Println("Escolha inválida. Escolhendo Português.")
		language = "pt"
	}
}

func playGame() {
	for {
		answer := generateAnswer()
		attempts := 0
		errors := 0

		fmt.Println(getLocalizedString("guess_number"))

		for {
			guess := getUserGuess()

			attempts++

			if guess < answer {
				fmt.Println(getLocalizedString("higher"))
			} else if guess > answer {
				fmt.Println(getLocalizedString("lower"))
			} else {
				fmt.Printf(getLocalizedString("congrats"), attempts)
				break
			}

			if guess != answer {
				errors++
				if errors == maxErrorsBeforeHint {
					if answer%2 == 0 {
						fmt.Println(getLocalizedString("hint_divisible_by_2"))
					} else if answer%3 == 0 {
						fmt.Println(getLocalizedString("hint_multiple_of_3"))
					} else {
						fmt.Println(getLocalizedString("hint_not_divisible_by_2_or_multiple_of_3"))
					}
				}
			}
		}
		fmt.Println("")
		fmt.Println(getLocalizedString("play_again"))
		if !playAgain() {
			break
		}
	}
}

func generateAnswer() int {
	return rand.Intn(100) + 1
}

func getUserGuess() int {
	var guess int
	_, err := fmt.Scan(&guess)
	if err != nil {
		fmt.Println(getLocalizedString("invalid_number"))
		return getUserGuess()
	}
	return guess
}

func playAgain() bool {
	var playAgain string
	fmt.Scan(&playAgain)
	fmt.Scanln() // Consumir o caractere de nova linha

	if language == "pt" {
		return playAgain == "s" || playAgain == "S"
	} else if language == "en" {
		return playAgain == "y" || playAgain == "Y"
	} else {
		return false
	}
}

func getLocalizedString(key string) string {
	switch language {
	case "en":
		// English strings
		switch key {
		case "guess_number":
			return "Guess the number between 1 and 100:"
		case "higher":
			return "Higher!"
		case "lower":
			return "Lower!"
		case "congrats":
			return "Congratulations! You guessed it in %d attempts."
		case "hint_divisible_by_2":
			return "Hint: The number is divisible by 2."
		case "hint_multiple_of_3":
			return "Hint: The number is a multiple of 3."
		case "hint_not_divisible_by_2_or_multiple_of_3":
			return "Hint: The number is neither divisible by 2 nor a multiple of 3."
		case "invalid_number":
			return "Invalid number. Please try again."
		case "play_again":
			return "Play again? (y/n)"
		default:
			return ""
		}
	case "pt":
		// Strings em Português
		switch key {
		case "guess_number":
			return "Adivinhe o número entre 1 e 100:"
		case "higher":
			return "Maior!"
		case "lower":
			return "Menor!"
		case "congrats":
			return "Parabéns! Você acertou em %d tentativas."
		case "hint_divisible_by_2":
			return "Dica: O número é divisível por 2."
		case "hint_multiple_of_3":
			return "Dica: O número é múltiplo de 3."
		case "hint_not_divisible_by_2_or_multiple_of_3":
			return "Dica: O número não é divisível nem por 2 nem é múltiplo de 3."
		case "invalid_number":
			return "Número inválido. Tente novamente."
		case "play_again":
			return "Jogar novamente? (s/n)"
		default:
			return ""
		}
	default:
		return ""
	}
}
