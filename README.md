# XOR Files with AVX2

This program performs a bitwise XOR operation on the contents of two files using AVX2 instructions for efficient processing. It's written in Go, utilizing cgo to integrate AVX2-based operations written in C for enhanced performance.

## Prerequisites

- Go (1.21 or later recommended)
- GCC compiler (for cgo to compile C code)

## Installation

1. Clone the repository to your local machine:

```bash
git clone https://github.com/qnfm/xor.git
```

2. Navigate to the cloned repository:

```bash
cd xor
```

3. Compile the program:

```bash
go build -o xor
```

## Usage

```bash
xor <input file 1> <input file 2> <output file>
```

Replace <input file 1>, <input file 2>, and <output file> with the paths to your input files and the desired output file path, respectively. Note that the input files must be of the same size.

## Benchmark

```bash
go test -benchmem -bench=.
```

## How It Works

The program reads the contents of two files into memory, performs a bitwise XOR operation between each corresponding byte in the files using AVX2 instructions, and writes the result to an output file. This is achieved by embedding C code in the Go program using cgo, allowing for the use of AVX2 instructions for the XOR operation.

## Contributing

Contributions are welcome! If you have suggestions for improvements or encounter any issues, please open an issue or submit a pull request.