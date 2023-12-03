package main

import (
	_ "embed"
	"log/slog"
	"slices"
	"strconv"
	"strings"
	"unicode"
)

//go:embed input.txt
var schematic string

type NumberCoord struct {
	Number int
	StartX int
	EndX   int
	Y      int
}

type SymbolCoord struct {
	Symbol rune
	X      int
	Y      int
}

func main() {
	var sum1, sum2 int
	var numCoords []NumberCoord
	var symbolCoords []SymbolCoord

	splittedSchematic := strings.Split(schematic, "\n")

	for y := 0; y < len(splittedSchematic); y++ {
		for x := 0; x < len(splittedSchematic[y]); x++ {
			r := rune(splittedSchematic[y][x])

			if r == '.' {
				continue
			}

			if unicode.IsDigit(r) {
				numIx := slices.IndexFunc(numCoords, func(c NumberCoord) bool {
					return c.EndX == x-1 && c.Y == y
				})
				num, _ := strconv.Atoi(string(r))

				if numIx > -1 {
					numCoords[numIx].EndX = x
					numCoords[numIx].Number = 10*numCoords[numIx].Number + num
				} else {
					numCoords = append(numCoords, NumberCoord{
						Number: num,
						StartX: x,
						EndX:   x,
						Y:      y,
					})
				}
			} else {
				symbolCoords = append(symbolCoords, SymbolCoord{
					Symbol: r,
					X:      x,
					Y:      y,
				})
			}
		}
	}

	for _, nc := range numCoords {
		symbols := filter(symbolCoords, func(sc SymbolCoord) bool {
			return (sc.Y == nc.Y ||
				sc.Y == nc.Y-1 ||
				sc.Y == nc.Y+1) &&
				sc.X >= nc.StartX-1 &&
				sc.X <= nc.EndX+1
		})

		if len(symbols) > 0 {
			sum1 += nc.Number
		}
	}

	for _, sc := range filter(symbolCoords, func(sc SymbolCoord) bool {
		return sc.Symbol == '*'
	}) {
		numbers := filter(numCoords, func(nc NumberCoord) bool {
			return (nc.Y == sc.Y ||
				nc.Y == sc.Y-1 ||
				nc.Y == sc.Y+1) &&
				sc.X >= nc.StartX-1 &&
				sc.X <= nc.EndX+1
		})

		if len(numbers) == 2 {
			sum2 += numbers[0].Number * numbers[1].Number
		}
	}

	slog.Default().Info("sum of all of the part numbers in the engine schematic", slog.Int("part", 1), slog.Int("sum", sum1))
	slog.Default().Info("sum of all of the gear ratios in the engine schematic", slog.Int("part", 2), slog.Int("sum", sum2))
}

func filter[T any](x []T, fn func(T) bool) []T {
	var fs []T

	for _, s := range x {
		if fn(s) {
			fs = append(fs, s)
		}
	}

	return fs
}
