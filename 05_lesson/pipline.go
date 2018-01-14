package main

import "fmt"

func main() {
	natural := make(chan int)
	squares := make(chan int)
	go counter(natural)
	go squarer(natural, squares)
	printer(squares)
}

func counter(out chan<- int) {
	for i := 0; i < 10; i++ {
		out <- i
		fmt.Println("Counter ", i)
	}
	close(out)
}

func squarer(in <-chan int, out chan<- int) {
	for v := range in {
		out <- v * v
		fmt.Println("Squarer ", v)
	}
	close(out)
}

func printer(in <-chan int) {
	for v := range in {
		fmt.Println(v)
	}
}
