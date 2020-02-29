package leetcode

import (
	"fmt"
	"testing"
)

//[[1],[2],[3],[]]
func TestCanVisitAllRooms1(t *testing.T) {
	fmt.Println(canVisitAllRooms([][]int{[]int{1}, []int{2}, []int{3}, []int{}}))
}

func TestCanVisitAllRooms2(t *testing.T) {
	fmt.Println(canVisitAllRooms([][]int{[]int{1, 3}, []int{3, 0, 1}, []int{2}, []int{0}}))
}

//[[],[1,1],[2,2]]
func TestCanVisitAllRooms3(t *testing.T) {
	fmt.Println(canVisitAllRooms([][]int{[]int{}, []int{1, 1}, []int{2, 2}})) //false
}

//[[]]
func TestCanVisitAllRooms4(t *testing.T) {
	fmt.Println(canVisitAllRooms([][]int{[]int{}})) //true
}
