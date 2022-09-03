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

// Q3. 掛け算 Multiply と 割り算 Divide についてもQ2 を同じ順番で実装してみましょう
//     一つのテストの中に複数回の関数の呼び出しと結果の確認のコードを書いてください。
