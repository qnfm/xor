package main

/*
#cgo CFLAGS: -mavx2
#include <immintrin.h>
#include <stdint.h>

// Function to XOR blocks of data using AVX2
void xor_avx2(uint8_t *dest, const uint8_t *src1, const uint8_t *src2, size_t len) {
    for (size_t i = 0; i < len; i += 32) {
        __m256i a = _mm256_loadu_si256((__m256i *)(src1 + i));
        __m256i b = _mm256_loadu_si256((__m256i *)(src2 + i));
        __m256i result = _mm256_xor_si256(a, b);
        _mm256_storeu_si256((__m256i *)(dest + i), result);
    }
}
*/
import "C"

import (
	"fmt"
	"log"
	"os"
	"unsafe"
)

// XorFiles reads the content of two files, performs XOR operation on them, and writes the result to an output file.
func XorFiles(file1, file2, output []byte) {

	// Call the C function to perform the XOR operation
	C.xor_avx2((*C.uint8_t)(unsafe.Pointer(&output[0])), (*C.uint8_t)(unsafe.Pointer(&file1[0])), (*C.uint8_t)(unsafe.Pointer(&file2[0])), C.size_t(len(file1)))

}

func main() {
	if len(os.Args) != 4 {
		fmt.Println("Usage: xor <input file 1> <input file 2> <output file>")
		return
	}

	// Retrieve file paths from command line arguments
	in1, err := os.ReadFile(os.Args[1])
	if err != nil {
		log.Println(err)
		return
	}
	in2, err := os.ReadFile(os.Args[2])
	if err != nil {
		log.Println(err)
		return
	}

	out := make([]byte, len(in1))
	// Perform the XOR operation
	XorFiles(in1, in2, out)

	outputPath := os.Args[3]
	err = os.WriteFile(outputPath, out, 0644)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println("XOR operation completed successfully.")
}
