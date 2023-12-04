package main

import (
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var input string

type Card struct {
	WinningNumbers []int
	ChosenNumbers  []int
}

func (c Card) Wins() int {
	var wins int

	for _, wn := range c.WinningNumbers {
		for _, n := range c.ChosenNumbers {
			if wn != n {
				continue
			}

			wins += 1
		}
	}

	return wins
}

func (c Card) Points() int {
	var points int

	wins := c.Wins()
	if wins > 0 {
		points += 1 << (wins - 1)
	}

	return points
}

func main() {
	splittedInput := strings.Split(input, "\n")
	cards := make([]Card, len(splittedInput))

	for i, si := range splittedInput {
		var card Card

		numbers := strings.Split(strings.Split(si, ": ")[1], " | ")

		winningNumbers := strings.Split(numbers[0], " ")
		for _, wn := range winningNumbers {
			if wn == "" {
				continue
			}

			pwn, _ := strconv.Atoi(wn)

			card.WinningNumbers = append(card.WinningNumbers, pwn)
		}

		chosenNumbers := strings.Split(numbers[1], " ")
		for _, cn := range chosenNumbers {
			if cn == "" {
				continue
			}

			pcn, _ := strconv.Atoi(cn)

			card.ChosenNumbers = append(card.ChosenNumbers, pcn)
		}

		cards[i] = card
	}

	var points int

	// Part 2: Points do not exist, but scratchcards do! AHHHHHHHHHHHHHHHHHHH
	// I suck at naming variables, so deal with it.
	cardCopies := len(cards)
	cardCount := make(map[int]int)

	for i, c := range cards {
		points += c.Points()

		cardCount[i] += 1

		for j := i + 1; j <= c.Wins()+i; j++ {
			cardCount[j] += cardCount[i]
			cardCopies += cardCount[i]
		}
	}

	slog.Default().Info("points worth in total", slog.Int("part", 1), slog.Int("points", points))
	slog.Default().Info("total scratchcards", slog.Int("part", 2), slog.Int("scratchcards", cardCopies))
}
