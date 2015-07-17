package main

import (
	"fmt"
	"io"
	"os"
)

// From http://pubs.opengroup.org/onlinepubs/9699919799/utilities/cksum.html
// polynomial has powers (32 and)
// 26 23 22 16 12 11 10 8 7 5 4 2 1 0
var P uint32 = (1 << 26) + (1 << 23) + (1 << 22) + (1 << 16) + (1 << 12) + (1 << 11) + (1 << 10) + (1 << 8) + (1 << 7) + (1 << 5) + (1 << 4) + (1 << 2) + (1 << 1) + (1 << 0)

func main() {
	var nBytes uint64
	var s uint32
	in := os.Stdin

	for {
		buffer := make([]byte, 1)
		count, err := in.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		nBytes += uint64(count)

		for i := 0; i < count; i += 1 {
			s = sum(P, s, buffer[i])
		}
	}

	// fmt.Println("raw", s)

	if nBytes == 0 {
		s = sum(P, s, 0)
	}
	for a := nBytes; a != 0; {
		var x byte = byte(nBytes & 0xff)
		s = sum(P, s, x)
		a >>= 8
	}
	for i := 0; i < 4; i += 1 {
		s = sum(P, s, 0)
	}
	fmt.Println(^s, nBytes)
}

// p: divisor polynomial (coefficient x^31 to X^0; coefficient
// x^32 is assumed to be 1).
// s: current sum
// b: additional byte
func sum(p uint32, s uint32, b byte) uint32 {
	for i := 0; i < 8; i += 1 {
		if s&(1<<31) != 0 {
			s = (s << 1) ^ p
		} else {
			s = (s << 1)
		}
		if b&0x80 != 0 {
			s ^= 1
		}
		b <<= 1
	}

	return s
}
