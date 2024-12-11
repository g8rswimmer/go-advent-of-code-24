package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("instructions.txt")
	handleError(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	handleError(err)

	p := &parser{
		mem: data,
		idx: 0,
	}
	ans := 0
	for {
		n1, n2, ok := p.Next()
		if !ok {
			break
		}
		ans += (n1 * n2)
	}
	fmt.Printf("uncorrupted mul instructions: %d\n", ans)

	p.idx = 0
	p.allowSkip = true
	part2 := 0
	for {
		n1, n2, ok := p.Next()
		if !ok {
			break
		}
		part2 += (n1 * n2)
	}
	fmt.Printf("uncorrupted mul instructions w/skip: %d\n", part2)
}

func handleError(err error) {
	if err == nil {
		return
	}
	panic(err.Error())
}
