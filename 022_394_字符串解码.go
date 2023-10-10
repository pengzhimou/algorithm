package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。

// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。

// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。

// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
// 示例 1：

// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：

// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：

// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：

// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"

func mainss1() {
	// aa := decodeString("3[a]12[bc]")
	// aa := decodeString("13[ac]")
	aa := decodeString("3[a2[c]]")
	fmt.Println(aa)

	aaa := "abc4]"
	fmt.Println(aaa[0])
	fmt.Println(aaa[3])
	fmt.Println(aaa[4])

}

func decodeString(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]

		fmt.Println("======", string(cur))

		if cur >= '0' && cur <= '9' {
			fmt.Println("======<10", string(cur))
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			fmt.Println("====== string and [", string(cur))

			stk = append(stk, string(cur))
			ptr++
		} else {
			fmt.Println("====== ]", string(cur))

			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" { // 逆向入站sub
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ { //反转sub
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1] //去掉[
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1] // 去掉repeat num
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	fmt.Println("geDIgits: ", ret)
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

// //////////////////////////////
var (
	src string
	ptr int
)

func decodeString1(s string) string {
	src = s
	ptr = 0
	return getString1()
}

func getString1() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigits1()
		ptr++
		str := getString1()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getString1()
}

func getDigits1() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}
