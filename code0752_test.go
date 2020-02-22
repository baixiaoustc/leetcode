package leetcode

import (
	"fmt"
	"testing"
)

func TestOpenLock1(t *testing.T) {
	fmt.Println(openLock([]string{"0201", "0101", "0102", "1212", "2002"}, "0202"))
}

func TestOpenLock2(t *testing.T) {
	fmt.Println(openLock([]string{"8888"}, "0009"))
}

func TestOpenLock3(t *testing.T) {
	fmt.Println(openLock([]string{"8887", "8889", "8878", "8898", "8788", "8988", "7888", "9888"}, "8888"))
}

func TestOpenLock4(t *testing.T) {
	fmt.Println(openLock([]string{"0000"}, "8888"))
}
