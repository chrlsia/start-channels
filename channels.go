package main

//https://golang.gr/2022/02/11/channels/
import "fmt"

func main() {
	//make a channel
	ch := make(chan int)

	//start a go function and through it send into the channel
	go func(x, y int) {
		result := x + y
		ch <- result
	}(12, 3)
	//receive from the channel
	number := <-ch
	//print the contents of the channel
	fmt.Printf("The number received from channel is: %d\n", number)

}

//next:range-channels
