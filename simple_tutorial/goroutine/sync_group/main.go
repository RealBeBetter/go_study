package main

func main() {
	done := false

	go func() {
		done = true
	}()

	for !done {
	}

	println("done!")
}
