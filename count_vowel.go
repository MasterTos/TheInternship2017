package main

import (
    "fmt"
    "os"
    "strings"
    "regexp"
)
    
func main() {
    str := os.Args[1]
    fmt.Println(countVowel(str))
}

func countVowel(str string) int {
    str = strings.ToLower(str)
    count := 0
    for _, ch := range str {
        if matched, _ := regexp.MatchString("[aeiou]", string(ch)); matched {
            count++
        }
    }
    return count
}