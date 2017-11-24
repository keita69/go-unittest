package main

import (
	"testing"
)

func Test_calc(t *testing.T) {
	var goods float64 = 1000
	var expected int = 1080
	result := calc(goods)
	if result != expected {
		t.Fatalf("計算が合いません。テスト値【%d】≠ 計算結果【%d】", expected, result)
	}
}
