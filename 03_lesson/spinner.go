package main

import (
	"fmt"
	"time"
)

func main() {
	go spinner(100 * time.Microsecond)
	const n = 45
	f := fibN(n)
	fmt.Printf("\rFibonacci(%d)= %d\n", n, f)
}

func fibN(n int) int {
	if n < 2 {
		return n
	}
	return fibN(n-1) + fibN(n-2)
}

func spinner(delay time.Duration) {
	for {
		for _, v := range `-\|/` {
			fmt.Printf("\r%c", v)
			time.Sleep(delay)
		}
	}
}

func fibN2(n int) int {
	f := fib()
	for i := 1; i < n; i++ {
		f()
	}
	return f()

}

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}
