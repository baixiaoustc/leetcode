package leetcode

import (
	"fmt"
	"testing"
)

func TestSpiralOrder1(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder2(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder3(t *testing.T) {
	m := [][]int{[]int{1}, []int{4}, []int{7}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder4(t *testing.T) {
	m := [][]int{[]int{1, 2, 3}, []int{4, 5, 6}, []int{7, 8, 9}, []int{0, 0, 0}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder5(t *testing.T) {
	m := [][]int{[]int{1, 2, 3, 4}, []int{5, 6, 7, 8}, []int{9, 0, 1, 2}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder6(t *testing.T) {
	m := [][]int{[]int{1}}
	fmt.Println(spiralOrder(m))
}

func TestSpiralOrder7(t *testing.T) {
	m := [][]int{[]int{}}
	fmt.Println(spiralOrder(m))
}

func TestString0(t *testing.T) {
	a := "abc"
	var b string
	b = "abc"
	if a == b {
		fmt.Println(a)
		var c []rune
		c = []rune(a)
		fmt.Println(string(c))
		c[0] = 'x'
		fmt.Println(string(c))
		c = append(c, 'y')
		fmt.Println(string(c))
	}
}
