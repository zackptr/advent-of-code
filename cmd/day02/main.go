package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var b []byte

type Game struct {
	Id   int
	Sets []GameSet
}

func (g Game) Power() int {
	var red, green, blue int

	for _, s := range g.Sets {
		red = max(red, s.Red)
		green = max(green, s.Green)
		blue = max(blue, s.Blue)
	}

	return red * green * blue
}

func (g Game) IsValid() bool {
	valid := true

	for _, s := range g.Sets {
		if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
			valid = false
		}
	}

	return valid
}

type GameSet struct {
	Red   int
	Green int
	Blue  int
}

func main() {
	var sumId, sumPower int

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		ug := strings.Split(s.Text(), ": ")
		id, _ := strconv.Atoi(strings.TrimPrefix(ug[0], "Game "))
		g := Game{
			Id: id,
		}

		for _, ugs := range strings.Split(ug[1], ";") {
			var gs GameSet

			for _, ucb := range strings.Split(ugs, ",") {
				sucb := strings.Split(strings.TrimSpace(ucb), " ")

				switch sucb[1] {
				case "red":
					red, _ := strconv.Atoi(sucb[0])
					gs.Red = red
				case "green":
					green, _ := strconv.Atoi(sucb[0])
					gs.Green = green
				case "blue":
					blue, _ := strconv.Atoi(sucb[0])
					gs.Blue = blue
				}
			}

			g.Sets = append(g.Sets, gs)
		}

		if g.IsValid() {
			sumId += g.Id
		}

		sumPower += g.Power()
	}

	slog.Default().Info("sum of the ids of the possible games", slog.Int("part", 1), slog.Int("sum", sumId))
	slog.Default().Info("sum of the power of the games", slog.Int("part", 2), slog.Int("sum", sumPower))
}
