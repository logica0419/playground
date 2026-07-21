package main

import "testing"

func BenchmarkMaximumLikelihood(b *testing.B) {
	received := ParseBitString(received)

	for b.Loop() {
		_ = MaximumLikelihood(received)
	}
}

func BenchmarkViterbiDecode(b *testing.B) {
	received := ParseBitString(received)

	for b.Loop() {
		_ = ViterbiDecode(received)
	}
}
