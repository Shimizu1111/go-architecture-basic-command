package main

import (
	"os"
	"testing"
)

func TestCalculatorAdd(t *testing.T) {
	os.Args = []string{"cmd", "add", "3", "4"}
	calculator(os.Args)

	// TODO: ここには出力内容の検証ロジックを追加
}

func TestCalculatorSubtract(t *testing.T) {
	os.Args = []string{"cmd", "sub", "7", "4"}
	calculator(os.Args)

	// TODO: ここには出力内容の検証ロジックを追加
}
