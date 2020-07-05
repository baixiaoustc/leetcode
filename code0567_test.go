package leetcode

import (
	"fmt"
	"testing"
)

func TestCheckInclusion1(t *testing.T) {
	fmt.Println(checkInclusion("ab", "eidbaooo"))
}

func TestCheckInclusion2(t *testing.T) {
	fmt.Println(checkInclusion("ab", "eidboaoo"))
}

func TestCheckInclusion3(t *testing.T) {
	fmt.Println(checkInclusion("abb", "eidbaboo"))
}

func TestCheckInclusion4(t *testing.T) {
	fmt.Println(checkInclusion("abb", "eidbaabb"))
}

func TestCheckInclusion5(t *testing.T) {
	fmt.Println(checkInclusion("abc", "eidbaxbb"))
}

func TestCheckInclusion6(t *testing.T) {
	fmt.Println(checkInclusion("trinitrophenylmethylnitramine", "dinitrophenylhydrazinetrinitrophenylmethylnitramine"))
}
