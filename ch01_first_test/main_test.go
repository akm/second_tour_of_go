package main

import (
	"testing"
)

// Q1. 足し算を行う関数Addを、同じディレクトリのファイル add.go に作成してください
func TestGolang(t *testing.T) {
	s := Add(1, 2)
	if s != 3 {
		t.FailNow()
	}
}
