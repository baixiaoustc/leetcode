package leetcode

import (
	"fmt"
	"testing"
)

func TestDailyTemperatures1(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}

func TestDailyTemperatures2(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73, 74, 75, 76}))
}

func TestDailyTemperatures3(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{99, 88, 77, 66}))
}

func TestDailyTemperatures4(t *testing.T) {
	fmt.Println(dailyTemperatures([]int{73}))
}
