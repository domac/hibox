package main

//dead lock examples
func main() {

	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		select {
		case c1 <- 1:
			println("write 1")
		case c2 <- 2:
			println("write 2")
		}
	}()

	go func() {
		select {
		case <-c1:
			println("read 1")
		case <-c2:
			println("read 2")
		}
	}()

	select {}
}
