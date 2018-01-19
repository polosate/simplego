package main

import (
	"fmt"
	"time"
)

func main() {
	f := newFoo()
	go f.first()
	go f.second()
	go f.third()

	time.Sleep(time.Millisecond * 100)
	fmt.Println(f.label)
}

type foo struct {
	label string
	ch1  chan bool
	ch2 chan bool
	ch3 chan bool
}

func newFoo() foo {
	return foo{
		label: "",
		ch1: make(chan bool),
		ch2: make(chan bool),
		ch3: make(chan bool),
	}
}

func (f *foo) first() {
	f.label += "1"

}

func (f *foo) second() {
	f.label += " 2"
}

func (f *foo) third() {
	f.label += " 3"
}
