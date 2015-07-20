package main

import (
	"fmt"
	"github.com/posec/cksum"
	"io"
	"os"
)

func main() {
	var nBytes uint64
	var s uint32
	in := os.Stdin

	for {
		buffer := make([]byte, 10000)
		count, err := in.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
			return
		}
		nBytes += uint64(count)

		s = sumBuffer(s, buffer[:count])

	}

	// fmt.Println("raw", s)

	if nBytes == 0 {
		s = cksum.Accumulate(s, 0)
	}
	for a := nBytes; a != 0; {
		var x byte = byte(a & 0xff)
		s = cksum.Accumulate(s, x)
		a >>= 8
	}
	//	for i := 0; i < 4; i += 1 {
	//		s = sum(P, s, 0)
	//	}
	fmt.Println(^s, nBytes)
}

func sumBuffer(s uint32, buffer []byte) uint32 {
	for i := 0; i < len(buffer); i += 1 {
		s = cksum.Accumulate(s, buffer[i])
	}
	return s
}
