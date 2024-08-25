package main

import (
	"testing"
)

// benchXOROperation performs the XOR operation on two byte slices of the given size.
func benchXOROperation(b *testing.B, size int) {
	// Generate two random byte slices of the specified size
	data1 := make([]byte, size)
	data2 := make([]byte, size)
	// Fill the slices with some data (skipped for brevity)

	output := make([]byte, size)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		// Call the C function to perform the XOR operation
		XorFiles(data1, data2, output)
	}
}

// Benchmark for 1KB
func BenchmarkXOR1KB(b *testing.B) {
	benchXOROperation(b, 1024) // 1KB
}

// Benchmark for 1MB
func BenchmarkXOR1MB(b *testing.B) {
	benchXOROperation(b, 1024*1024) // 1MB
}

// Benchmark for 1GB
func BenchmarkXOR1GB(b *testing.B) {
	benchXOROperation(b, 1024*1024*1024) // 1GB
}
