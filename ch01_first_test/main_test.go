package main

import (
	"testing"
)

func TestGolang(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.FailNow()
	}
}

func TestSubstract(t *testing.T) {
	s := Substract(6, 2)
	if s != 4 {
		t.FailNow()
	}
}

// Q3. 掛け算 Multiply と 割り算の商と剰余を求める関数 Divide についてもQ2 を同じ順番で実装してみましょう。
//     (商は英語で quotient, 剰余は remainder と呼ばれ、qとrと略すことが多いです)
//     一つのテストの中に複数回の関数の呼び出しと結果の確認のコードを書いてください。
func TestMultiply(t *testing.T) {
	if m := Multiply(6, 2); m != 12 {
		t.FailNow()
	}
	if m := Multiply(0, 2); m != 0 {
		t.FailNow()
	}
	if m := Multiply(100, 0); m != 0 {
		t.FailNow()
	}
}

func TestDivide(t *testing.T) {
	if q, r := Divide(6, 2); q != 3 || r != 0 {
		t.FailNow()
	}
	if q, r := Divide(7, 3); q != 2 || r != 1 {
		t.FailNow()
	}
}
