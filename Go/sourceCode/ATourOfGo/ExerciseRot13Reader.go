package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

const (
	ByteA = int('A')
	ByteZ = int('Z')
	Bytea = int('a')
	Bytez = int('z')
)

// a b c d e f
// 1 2 3 4 5 6
// e -> b
// (5 + 6/2) - 6 + 1 - 1 = 2
func convertByRot13(b *byte) {
	i := int(*b) + 13
	if i < ByteZ {
		*b = uint8(i)
	} else if i < Bytea {
		// fmt.Printf("Z %c %3d %c %3d\n", byte(*b), byte(*b), byte(i-ByteZ+ByteA), byte(i-ByteZ+ByteA-1))
		*b = byte(i - ByteZ + ByteA - 1)
	} else if i < Bytez {
		// fmt.Printf("a %c %3d %c %3d\n", byte(*b), byte(*b), byte(i), byte(i))
		*b = byte(i)
	} else {
		// fmt.Printf("z %c %3d %c %3d\n", byte(*b), byte(*b), byte(i-Bytez+Bytea), byte(i-Bytez+Bytea-1))
		*b = byte(i - Bytez + Bytea - 1)
	}
}

func (reader rot13Reader) Read(bytes []byte) (int, error) {
	n, err := reader.r.Read(bytes)

	for i := range bytes {
		convertByRot13(&bytes[i])
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
	fmt.Printf("\n")
}
