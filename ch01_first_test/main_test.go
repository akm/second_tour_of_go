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

// Q2. 引き算を行う関数 Substract を substruct.go に実装します。
// 1. テストTestSubstractを作成します（コンパイルエラー)
// 2. 必ず0を返す関数Substractを作成（コンパイルエラーなし）
// 3. テストを実行します（テスト失敗）
// 4. Substractを実装します（テスト成功）
func TestSubstract(t *testing.T) {
	s := Substract(6, 2)
	if s != 4 {
		t.FailNow()
	}
}
