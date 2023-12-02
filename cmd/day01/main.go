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
	var sum1, sum2 int

	s := bufio.NewScanner(bytes.NewReader(b))
	for s.Scan() {
		sum1 += calibration(s.Text())

		t := s.Text()
		for i, digit := range digits {
			t = strings.ReplaceAll(t, digit, fmt.Sprintf("%s%d%s", digit[0:1], i+1, digit[len(digit)-1:]))
		}

		sum2 += calibration(t)
	}

	slog.Default().Info("sum of all of the calibration values", slog.Int("part", 1), slog.Int("sum", sum1))
	slog.Default().Info("sum of all of the calibration values", slog.Int("part", 2), slog.Int("sum", sum2))
}

func calibration(str string) int {
	var nums []int

	for _, sp := range strings.Split(str, "") {
		num, err := strconv.Atoi(sp)
		if err != nil {
			continue
		}

		nums = append(nums, num)
	}

	if len(nums) == 0 {
		return 0
	}

	return nums[0]*10 + nums[len(nums)-1]
}
