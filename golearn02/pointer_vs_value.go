package main

import "fmt"

type ByteSlice []byte

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	*p = slice
	return len(data), nil
}

func main() {
	var b ByteSlice
	a := []byte{1}
	b.Write(a)
	fmt.Fprintf(&b, "This hour has %d days\n", 7)
	fmt.Printf("%T---%v---%d",b,b,b)
}
