package main

import "fmt"

func main() {

	a := &A{}
	a.n = 10

	b := &B{}
	b.A.n = 20
	b.n = 30

	fmt.Println(b.A.n)
	fmt.Println(b.n)
	fmt.Println(a.n)

	b.test()
	b.A.test()
}

type A struct {
	n int32
}

type B struct {
	A
	n int32
}

func (a *A) test() {
	fmt.Println("A...", *a)
}

func (b *B) test() {
	fmt.Println("B...", *b)
}
