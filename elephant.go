package main

import (
	"fmt"
	"math"
)

func main() {
	var l, r uint64
	fmt.Scan(&l)
	fmt.Scan(&r)
	fmt.Print(f(r) - f(l-1))
}

func f(x uint64) uint64 {
	if x < 10 {
		return x
	}
	d := countDigits(x)
	var sum int = 9
	for i := 0; i+2 < d; i++ {
		sum += 9 * int(math.Pow10(i))
	}
	var l, r int
	x, l = lastDigit(x, d)
	r, x = int(x%10), x/10
	var temp int = 0
	for dt := d - 2; dt > 0; dt-- {
		var lt int
		x, lt = lastDigit(x, dt)
		if lt != 0 {
			temp += lt * int(math.Pow10(dt-1))
		}
	}
	sum += (l-1)*int(math.Pow10(d-2)) + temp
	if r < l {
		return uint64(sum)
	}
	return uint64(sum + 1)
}

func lastDigit(n uint64, d int) (uint64, int) {
	p := uint64(math.Pow10(d - 1))
	l := 0
	for ; n >= p; l++ {
		n -= p
	}
	return n, l
}

func countDigits(n uint64) int {
	var count int = 0
	for n > 0 {
		n /= 10
		count++
	}
	return count
}
