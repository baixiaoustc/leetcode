package leetcode

import (
	"fmt"
	"testing"
)

func TestFindMaxConsecutiveOnes1(t *testing.T) {
	a := []int{1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes2(t *testing.T) {
	a := []int{1, 1, 1, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes3(t *testing.T) {
	a := []int{0, 0, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes4(t *testing.T) {
	a := []int{1, 1, 0, 1, 1, 1, 0}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes5(t *testing.T) {
	a := []int{1, 0, 1, 0, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes6(t *testing.T) {
	a := []int{1, 1, 1, 1, 1, 1}
	fmt.Println(findMaxConsecutiveOnes(a))
}

func TestFindMaxConsecutiveOnes7(t *testing.T) {
	a := []int{0, 0}
	fmt.Println(findMaxConsecutiveOnes(a))
}
