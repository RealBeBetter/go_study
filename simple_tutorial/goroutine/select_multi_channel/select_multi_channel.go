package main

func main() {
	selectFromMultiChan()
}

func selectFromMultiChan() {
	stringChan := make(chan string)
	intChan := make(chan int)

	go func() {
		stringChan <- ""
	}()
	go func() {
		intChan <- 1
	}()

	select {
	case msg := <-stringChan:
		println("received string:", msg)
	case number := <-intChan:
		println("received int:", number)
	}

	println("done")
}
