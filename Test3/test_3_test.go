package test3

import (
	"fmt"
	"testing"
)

type Bebek struct {
	energi       int
	hidup        bool
	bisaTerbang  bool
	suaraTerbang string
}

func Mati(b *Bebek) {
	b.hidup = false
}
func Terbang(b *Bebek) {
	if b.energi > 0 && b.hidup == true && b.bisaTerbang {
		fmt.Println(b.suaraTerbang)
		b.energi -= 1
		if b.energi == 0 {
			Mati(b)
		}
	}
}
func Makan(b *Bebek) {
	if b.energi > 0 && b.hidup == true {
		b.energi += 1
	}
}

func TestTerbang(t *testing.T) {
	b := &Bebek{
		energi:       5,
		hidup:        true,
		bisaTerbang:  true,
		suaraTerbang: "terbang",
	}
	for i := 0; i < 100; i++ {
		Terbang(b)
	}
}
