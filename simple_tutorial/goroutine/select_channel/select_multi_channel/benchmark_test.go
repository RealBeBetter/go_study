package main

import "testing"

func BenchmarkSelectFromMultiChan(b *testing.B) {
	for i := 0; i < b.N; i++ {
		selectFromMultiChan()
	}
}
