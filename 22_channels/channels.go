package main

import "fmt"

// send
// func processNum(numChan chan int) {
// 	for num := range numChan {
// 		fmt.Println("Processing number", num)
// 		time.Sleep(time.Second)
// 	}
// }

// receive
// func sum(result chan int, num1 int, num2 int) {
// 	ans := num1 + num2
// 	result <- ans
// }

// goroutine synchronizer
// func task(done chan bool) {
// 	defer func() { done <- true }()
// 	fmt.Println("process...")
// }

// func mailer(emailchan <-chan string, done chan<- bool) {
// 	defer func() { done <- true }()
// 	for email := range emailchan {
// 		fmt.Println("Sending mail to:", email)
// 	}
// }

func main() {
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 10
	}()

	go func() {
		chan2 <- "pong"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Val := <-chan1:
			fmt.Println("Received data from chan1:", chan1Val)
		case chan2Val := <-chan2:
			fmt.Println("Received data from chan2:", chan2Val)
		}
	}

	// emailChan := make(chan string, 100) // Buffered channel
	// done := make(chan bool)

	// go mailer(emailChan, done)

	// for i := 0; i < 100; i++ {
	// 	emailChan <- fmt.Sprintf("%d@example.in", i)
	// }

	// // closing emailChan is important
	// close(emailChan)
	// <-done

	// done := make(chan bool)
	// go task(done)
	// <-done // blocks until we get value in done channel from task func.

	// result := make(chan int)
	// go sum(result, 10, 20)
	// ans := <-result // blocking when channel is empty. This is the reason why we need not use time.Sleep in this case as <-result is blocking until there is data in the channel which will come in only after the sum() function is executed
	// fmt.Println(ans)

	// numChan := make(chan int)

	// go processNum(numChan)

	// for {
	// 	numChan <- rand.IntN(100)
	// }

	// messageChan := make(chan string)

	// messageChan <- "pink"

	// msg := <-messageChan

	// fmt.Println(msg)
}
