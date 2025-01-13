package main

import (
	"bufio"
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	Title   string
	Options []string
	Answer  int
}
type GameState struct {
	Name      string
	Points    int
	Questions []Question
}

func (game *GameState) ProcessCSV() {
	file, err := os.Open("quiz-go.csv")

	if err != nil {
		panic("erro ao ler o arquivo")
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		panic("Erro ao ler csv")
	}

	for index, record := range records {
		//fmt.Println(index, record)
		if index > 0 {
			correctAnswer, _ := toInt(record[5])
			question := Question{
				Title:   record[0],
				Options: record[1:5],
				Answer:  correctAnswer,
			}

			game.Questions = append(game.Questions, question)

		}
	}

}

func toInt(s string) (int, error) {
	i, err := strconv.Atoi(s)

	if err != nil {
		return 0, errors.New("não é permitido um caractere diferente, por favor, digite um número")
	}

	return i, nil
}

func (game *GameState) Init() {
	fmt.Println("Seja bem vindo(a) ao quiz")
	fmt.Print("Escreva o seu nome: ")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		panic("Erro ao ler a string")
	}
	fmt.Println("---------------------------")
	game.Name = name
	fmt.Printf("Vamos ao jogo %s\n!!!", strings.ToUpper(game.Name))
}
func (game *GameState) Run() {
	for index, question := range game.Questions {
		fmt.Printf("\033[33m %d. %s \033[0m\n", index+1, question.Title)

		for j, option := range question.Options {
			fmt.Printf("[%d] %s\n", j+1, option)
		}

		fmt.Print("Digite uma alternativa: ")

		var answer int
		var err error

		for {
			reader := bufio.NewReader(os.Stdin)
			read, _ := reader.ReadString('\n')

			answer, err = toInt(read[:len(read)-1])

			if err != nil {
				fmt.Println(err.Error())
				continue
			}
			break
		}

		fmt.Printf("\nResposta Escolhida: %d\n\n", answer)

		if answer == question.Answer {
			fmt.Println("Resposta correta!")
			game.Points += 10
		} else {
			fmt.Println("Resposta incorreta!")
			fmt.Println("------------------------")
		}
	}
}

func main() {
	game := &GameState{}

	go game.ProcessCSV()
	game.Init()
	game.Run()

	fmt.Printf("\nFim de Jogo, você fez %d pontos.\n", game.Points)
}
