package math

import "testing"

const testNumber = 12345678910

func BenchmarkPrimeFactor(b *testing.B) {
	for b.Loop() {
		PrimeFactor(testNumber)
	}
}

func BenchmarkPrimeFactor1(b *testing.B) {
	for b.Loop() {
		PrimeFactor1(testNumber)
	}
}
