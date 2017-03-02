package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	input := os.Args[1]
	content, err := ioutil.ReadFile(input)
	if err != nil {
		panic(err)
	}
	line := strings.Split(strings.ToLower(string(content)), "\n")
	count := 0
	for _, str := range line {
		if isPalindome(str) {
			count++
		}
	}
	fmt.Println(count)
}

func isPalindome(str string) bool {
	k := len(str) - 1
	for i := 0; i < k; i++ {
		for !isAlphabet(str[i]) {
			i++
		}
		for !isAlphabet(str[k]) {
			k--
		}
		if str[i] != str[k] {
			return false
		}
		k--
	}
	return true
}

func isAlphabet(ch byte) bool {
	if ch >= 'a' && ch <= 'z' {
		return true
	}
	return false
}
