package leetcode

import (
	"fmt"
	"testing"
)

func TestCanPartition1(t *testing.T) {
	fmt.Println(canPartitionZeroOne([]int{1, 5, 11, 5}))
}

func TestCanPartition2(t *testing.T) {
	fmt.Println(canPartitionZeroOne([]int{1, 2, 3, 5}))
}

//[1,2,5]
func TestCanPartition3(t *testing.T) {
	fmt.Println(canPartitionZeroOne([]int{1, 2, 5}))
}

//[1,1,1,1,1,1,2,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,1,100,99,100]
func TestCanPartition4(t *testing.T) {
	fmt.Println(canPartitionZeroOne([]int{1, 1, 1, 1, 1, 1, 2, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 100, 99, 100}))
}

func TestKnapsackZeroOne2DP1(t *testing.T) {
	fmt.Println(knapsackZeroOne1DP([]int{2, 1, 3}, []int{4, 2, 3}, 4))
}

func TestKnapsackZeroOne2DP2(t *testing.T) {
	fmt.Println(knapsackZeroOne1DP([]int{2, 3, 4, 5}, []int{3, 4, 5, 6}, 8))
}

func TestKnapsackZeroOne2DP3(t *testing.T) {
	fmt.Println(knapsackZeroOne1DP([]int{1, 2, 3}, []int{6, 10, 12}, 5))
}
