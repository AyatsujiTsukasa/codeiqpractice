package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	for sc.Scan() {
		input := sc.Text()
		output := solve(input)
		fmt.Println(output)
	}
}

func solve(input string) string {
	i, _ := strconv.Atoi(input)
	if i%2 == 0 {
		return "0"
	}
	ans := 1999999/(i*2) + 1999999/i%2
	return strconv.Itoa(ans)
}
