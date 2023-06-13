package main

import (
	"fmt"
	"math/rand"
	"time"
)

const maxErrorsBeforeHint = 2

var language string
var round int
var tentativasPorJogada []int
var somaTentativas int

func main() {
	rand.Seed(time.Now().UnixNano())
	chooseLanguage()
	playGame()
}

func chooseLanguage() {
	fmt.Println("Escolha o seu idioma (Choose a language/言語を選択してください):")
	fmt.Println("1. Português (Portuguese/ポルトガル語)")
	fmt.Println("2. English (Inglês/英語)")
	fmt.Println("3. 日本 (Japonês/Japanese)")
	var choice int
	fmt.Scan(&choice)
	fmt.Scanln() // Consumir o caractere de nova linha

	switch choice {
	case 1:
		language = "pt"
	case 2:
		language = "en"
	case 3:
		language = "jp"
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
		round := 0

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
				tentativasPorJogada = append(tentativasPorJogada, attempts)
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
			for _, at := range tentativasPorJogada {
				somaTentativas += at
				round++
				if language == "pt" {
					fmt.Printf("Você utilizou %v tentativas no round %v\n", at, round)
				} else if language == "en" {
					fmt.Printf("You used %v attempts in round %v\n", at, round)
				} else if language == "jp" {
					fmt.Printf("ラウンド %v で %v トライを使用しました\n", round, at)
				}
			}
			if language == "pt" {
				fmt.Printf("O total de tentativas foi: %v", somaTentativas)
			} else if language == "en" {
				fmt.Printf("The total number of attempts is: %v", somaTentativas)
			} else if language == "jp" {
				fmt.Printf("合計試行回数は %v 回です", somaTentativas)
			}
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
	} else if language == "jp" {
		return playAgain == "h" || playAgain == "H"
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
	case "jp":
		// Japanese strings
		switch key {
		case "guess_number":
			return "1から100までの数字を推測してください:"
		case "higher":
			return "もっと大きいです!"
		case "lower":
			return "もっと小さいです!"
		case "congrats":
			return "おめでとうございます！%d回で正解しました。"
		case "hint_divisible_by_2":
			return "ヒント：その数字は2で割り切れます。"
		case "hint_multiple_of_3":
			return "ヒント：その数字は3の倍数です。"
		case "hint_not_divisible_by_2_or_multiple_of_3":
			return "ヒント：その数字は2で割り切れず、かつ3の倍数でもありません。"
		case "invalid_number":
			return "無効な数字です。もう一度やり直してください。"
		case "play_again":
			return "もう一度プレイしますか？ (h/i)"
		default:
			return ""
		}
	default:
		return ""
	}
}
