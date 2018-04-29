package main

import (
	"testing"
)

func TestRepeatChar(t *testing.T) {
	res := RepeatChar(6, 'q')
	expect := "qqqqqq"
	if res != expect {
		t.Fatal("Ridi")
	}
}

func BenchmarkRepeatChar(b *testing.B) {
	RepeatChar(b.N, 'w')
}
