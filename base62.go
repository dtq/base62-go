package main

import (
	"fmt"
	"math/big"
	"time"
)

var b62Alph = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var b62len = big.NewInt(62)
var runeMap = make(map[rune]int)

func init() {
	// Initialize the map of each rune to it's index value
	for i, c := range b62Alph {
		runeMap[c] = i
	}
}

func b62encode(num *big.Int) string {
	var encoding []rune
	rem := big.NewInt(0)
	zero := big.NewInt(0)

	for num.Cmp(zero) > 0 {
		num.DivMod(num, b62len, rem)
		r := b62Alph[rem.Int64()]
		encoding = append([]rune{r}, encoding...)
	}
	return string(encoding)

}

func divmod(a *big.Int, b *big.Int) (*big.Int, *big.Int) {
	return a.DivMod(a, b, nil)
}

func main() {
	start := time.Now()
	var b62 string

	num := new(big.Int)

	for i := 0; i < 1000000; i++ {
		fmt.Sscan("280997858289516668860144855962706392315", num)
		b62 = b62encode(num)
	}

	fmt.Println(time.Since(start))
	fmt.Println(b62)
}
