package main

import "fmt"

func lengthOfNonRepeatingSubStr(s string) int{
	lastOccurred := make(map[rune]int)
	start := 0
	maxLenth := 0
	for i, ch := range []rune(s){
		if lastI, ok := lastOccurred[ch];lastI >= start && ok{
			start = lastI + 1
		}
		if i - start + 1 > maxLenth{
			maxLenth = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLenth
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcabcbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("bbbbbbb"))
	fmt.Println(lengthOfNonRepeatingSubStr("黑化肥挥发发灰会飞回灰化肥挥发发黑会飞花"))
}
