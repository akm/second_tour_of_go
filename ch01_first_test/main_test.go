package main

import (
	"testing"
)

func TestGolang(t *testing.T) {
	// Suggest テストに失敗するようにパラメータや期待する結果を変更して、テストを実行してみてください。
	s := Add(1, 2)
	if s != 3 {
		t.FailNow()
	}
}
