package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for sc.Scan() {
		input := sc.Text()
		out := convert(input)
		fmt.Println(out)
	}
}

func convert(input string) string {
	output := []rune(input)
	for {
		i := findPattern(output)
		if i == -1 {
			break
		} else {
			output = delete(output, i)
			output = delete(output, i)
		}
	}
	return string(output)
}

func findPattern(input []rune) int {
	for i := 0; i < len(input)-1; i++ {
		if int(input[i]) == int(input[i+1])+1 || int(input[i]) == int(input[i+1])-1 {
			return i
		}
	}
	return -1
}

func delete(s []rune, i int) []rune {
	s = append(s[:i], s[i+1:]...)
	n := make([]rune, len(s))
	copy(n, s)
	return n
}
