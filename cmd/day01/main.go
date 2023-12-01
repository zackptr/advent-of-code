package main

import (
	"bufio"
	"bytes"
	_ "embed"
	"fmt"
	"log/slog"
	"strconv"
	"strings"
)

//go:embed input.txt
var b []byte

var digits = []string{"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	var (
		sum1 int
		sum2 int
	)

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		sum1 += calibration(s.Text())

		t := s.Text()
		for i, digit := range digits {
			t = strings.ReplaceAll(t, digit, fmt.Sprintf("%v%v%v", digit[0:1], i+1, digit[len(digit)-1:]))
		}

		sum2 += calibration(t)
	}

	slog.Default().With(slog.Int("part", 1)).Info("sum of all of the calibration values", slog.Int("sum", sum1))
	slog.Default().With(slog.Int("part", 2)).Info("sum of all of the calibration values", slog.Int("sum", sum2))
}

func calibration(str string) int {
	var nums []int

	split := strings.Split(str, "")
	for _, sp := range split {
		num, err := strconv.Atoi(sp)
		if err != nil {
			continue
		}

		nums = append(nums, num)
	}

	if len(nums) < 1 {
		return 0
	}

	if len(nums) < 2 {
		nums = append(nums, nums[0])
	}

	val, _ := strconv.Atoi(fmt.Sprintf("%d%d", nums[0], nums[len(nums)-1]))

	return val
}
