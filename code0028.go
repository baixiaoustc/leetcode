/*
实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/implement-strstr
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

package leetcode

func strStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	for i := 0; i < len(haystack); i++ {

		n := 0
		for j := range needle {
			if i+n >= len(haystack) {
				return -1
			}

			if haystack[i+n] != needle[j] {
				break
			}
			n++
		}

		if n == len(needle) {
			return i
		}
	}

	return -1
}

//sunday algorithm
func strStrSunday(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)
	if m > n {
		return -1
	}

	shift := make(map[int32]int)
	for i, c := range needle {
		shift[c] = m - i
	}

	for i := 0; i <= n-m; {
		tmp := haystack[i : i+m]
		if tmp == needle {
			return i
		}

		if i+m+1 > n {
			return -1
		}

		next := haystack[i+m : i+m+1]
		var a rune
		a = rune(next[0])
		if v, ok := shift[a]; ok {
			i += v
		} else {
			i += m + 1
		}
	}

	return -1
}

//kmp algorithm
//https://leetcode-cn.com/problems/implement-strstr/solution/kmp-suan-fa-xiang-jie-by-labuladong/
func strStrKMP(haystack string, needle string) int {
	if needle == "" {
		return 0
	}

	n := len(haystack)
	m := len(needle)
	if m > n {
		return -1
	}

	txt := make([]rune, n)
	pat := make([]rune, m)

	for i, c := range haystack {
		txt[i] = c
	}
	for i, c := range needle {
		pat[i] = c
	}

	//create dp
	dp := make([][256]int32, m)
	dp[0][pat[0]] = 1

	var x int32
	for j := 1; j < m; j++ {
		for c := 0; c < 256; c++ {
			dp[j][c] = dp[x][c]
		}
		dp[j][pat[j]] = int32(j + 1)
		x = dp[x][j]
	}

	//search
	var j int32 //initial state
	for i := 0; i < n; i++ {
		j = dp[j][txt[i]]
		if j == int32(m) {
			return i - m + 1
		}
	}

	return -1
}
