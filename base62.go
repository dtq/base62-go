package base62

import (
	"errors"
	"fmt"
	"math/big"
)

var b62Alph = []rune("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
var b62len = intToBig(len(b62Alph))
var runeMap = make(map[rune]*big.Int)

func init() {
	// Initialize the map of each rune to its index value
	for i, c := range b62Alph {
		runeMap[c] = intToBig(i)
	}
}

// EncodeBig takes a big.Int and returns a base62 encoded string
func EncodeBig(num *big.Int) (senc string, err error) {
	var encoding []rune
	rslt := new(big.Int)
	rslt.Set(num)
	rem := big.NewInt(0)
	zero := big.NewInt(0)

	if num.Cmp(zero) == 0 { // equals zero
		return string(b62Alph[0]), nil
	}
	if num.Cmp(zero) == -1 { // is negative
		return "", fmt.Errorf("base62: negative number %q", num)
	}

	for rslt.Cmp(zero) > 0 {
		rslt.DivMod(rslt, b62len, rem)
		r := b62Alph[rem.Int64()]
		encoding = append([]rune{r}, encoding...)
	}
	return string(encoding), nil
}

// EncodeStr takes a string representation of an integer (like "150") and returns
// a base62 encoded string
func EncodeStr(s string) (senc string, err error) {
	if s == "" {
		return "", errors.New("base62: empty string")
	}
	num := strToBig(s)
	return EncodeBig(num)
}

func Decode(s string) (num *big.Int, err error) {
	num = new(big.Int)
	if s == "" {
		return num, errors.New("base62: empty string")
	}

	for _, c := range s {
		num.Mul(num, b62len)
		num.Add(num, runeMap[c])
	}

	return num, nil
}

// strToBig takes a string representation of an integer (like "123") and returns a pointer
// to big.Int
func strToBig(s string) *big.Int {
	num := new(big.Int)
	fmt.Sscan(s, num)
	return num
}

// intToBig takes an int and returns a pointer to big.Int
func intToBig(num int) *big.Int {
	return big.NewInt(int64(num))
}
