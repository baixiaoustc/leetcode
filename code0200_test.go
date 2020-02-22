package leetcode

import (
	"fmt"
	"testing"
)

func TestNumIslands(t *testing.T) {
	grid := [][]byte{[]byte{'1', '1', '1', '1', '0'}, []byte{'1', '1', '0', '1', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '0', '0', '0'}}
	fmt.Println(numIslandsUnion(grid))
}

func TestNumIslands1(t *testing.T) {
	grid := [][]byte{[]byte{'1', '1', '0', '0', '0'}, []byte{'1', '1', '0', '0', '0'}, []byte{'0', '0', '1', '0', '0'}, []byte{'0', '0', '0', '1', '1'}}
	fmt.Println(numIslandsUnion(grid))
}

func TestByte(t *testing.T) {
	if 1 == byte(1) {
		fmt.Println("match")
	}
}
