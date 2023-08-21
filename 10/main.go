package main

import (
	"fmt"
)

const input string = "1113122113"

func main() {
	r := rbuv{[]byte(input), 0}
	fmt.Println("Part 1:", applyN(&r, 40))
	r = rbuv{[]byte(input), 0}
	fmt.Println("Part 2:", applyN(&r, 50))
}

func applyN(input *rbuv, times int) int {
	for i := 0; i < times; i++ {
		w := wbuv{[]byte{}}
		lookAndSaySequence(input, &w)
		input = &rbuv{w.data, 0}
	}
	return len(input.data)
}

type rbuv struct {
	data []byte
	pos  int
}

type wbuv struct {
	data []byte
}

func (b *rbuv) ReadByte() byte {
	if b.pos < len(b.data) {
		return b.data[b.pos]
	} else {
		return byte(0)
	}
}

func (b *rbuv) Advance() {
	b.pos++
}

func (b *wbuv) Write(i int, rb byte) {
	s := byte(i + 48)
	ap := []byte{s, rb}
	b.data = append(b.data, ap...)
}

func lookAndSaySequence(r *rbuv, w *wbuv) {
	prevRune := r.ReadByte()
	count := 0

	for r.pos <= len(r.data) {
		b := r.ReadByte()
		r.Advance()

		if b != prevRune {
			w.Write(count, r.data[r.pos-2])
			count = 0
		}
		count++
		prevRune = b
	}
}
