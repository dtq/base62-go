package base62

import (
	"math/big"
	"testing"
)

func TestEncodeBig(t *testing.T) {
	cases := []struct {
		in   *big.Int
		want string
	}{
		{intToBig(959248973882093004), "18rMy76iPka"},
		{intToBig(0), "0"},
		{intToBig(-5), ""},
	}
	for _, c := range cases {
		got, _ := EncodeBig(c.in)
		if got != c.want {
			t.Errorf("EncodeBig(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestEncodeStr(t *testing.T) {
	cases := []struct {
		in, want string
	}{
		{"95924897388209300414612745184451631672", "2CAqu3cYjyiuCgIsDjlZtY"},
		{"0", "0"},
		{"-5", ""},
		{"", ""},
	}
	for _, c := range cases {
		got, _ := EncodeStr(c.in)
		if got != c.want {
			t.Errorf("EncodeStr(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}

func TestDecode(t *testing.T) {
	cases := []struct {
		in   string
		want *big.Int
	}{
		{"2CAqu3cYjyiuCgIsDjlZtY", strToBig("95924897388209300414612745184451631672")},
		{"0", strToBig("0")},
		{"", strToBig("0")},
	}
	for _, c := range cases {
		got, _ := Decode(c.in)
		if got.Cmp(c.want) != 0 {
			t.Errorf("Decode(%q) == %q, want %q", c.in, got, c.want)
		}
	}
}
