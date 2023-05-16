package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(palindromo("Luz azul"))
}

func palindromo(word string) bool {
	fmt.Println(word)
	//word = strings.ToUpper(word)
	word = strings.ToLower(word)
	//word = strings.Replace(word, " ", "", 1)
	word = strings.ReplaceAll(word, " ", "")
	fmt.Println(word)
	invert := reverse(word)
	fmt.Println(invert)
	return invert == word
}

func reverse(word string) string {
	array := strings.Split(word, "")
	arrayInverse := make([]string, 0)

	for i := len(array) - 1; i >= 0; i-- {
		arrayInverse = append(arrayInverse, array[i])
	}
	// fmt.Println(array)
	// fmt.Println(arrayInverse)
	return strings.Join(arrayInverse, "")
}
