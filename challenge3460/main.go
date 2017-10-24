package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

var (
	sc     = bufio.NewScanner(os.Stdin)
	one    = big.NewInt(1)
	two    = big.NewInt(2)
	four   = big.NewInt(4)
	lookup = make(map[int]*big.Int, 2015)
	m      = big.NewInt(0)
	n      = big.NewInt(0)
)

func main() {
	for sc.Scan() {
		input := sc.Text()
		output := solve(input)
		fmt.Println(output)
	}
}

func solve(input string) string {
	inputArray := strings.Split(input, " ")
	intm, _ := strconv.ParseInt(inputArray[0], 10, 64)
	intn, _ := strconv.ParseInt(inputArray[1], 10, 64)
	m.SetInt64(intm)
	n.SetInt64(intn)
	ans := combination(new(big.Int).Mul(two, n), n)
	ans = ans.Mul(ans, big.NewInt(int64(math.Pow(float64(4), float64(intm-intn)))))
	ans = ans.Mul(ans, combination(m, n))
	return ans.String()
}

// 階乗を計算します
func factorial(n *big.Int) (result *big.Int) {
	if n.Cmp(one) < 1 {
		return one
	}

	result = new(big.Int).Set(n)
	result = new(big.Int).Mul(result, factorial(n.Sub(n, one)))

	lookup[int(n.Int64())] = result

	return
}

// 組み合わせ n C k を計算します
func combination(n, k *big.Int) *big.Int {

	a := factorial(new(big.Int).Set(k))
	b := factorial(new(big.Int).Sub(n, k))

	c := new(big.Int).Mul(a, b)
	d := factorial(new(big.Int).Set(n))

	return d.Div(d, c)
}
