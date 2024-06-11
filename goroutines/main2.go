package main

import (
	"fmt"
	//"time"
	//"time"
)


func main() {
	
	ch := make(chan int)
	completed := make (chan bool)
	go sender(ch, completed)
	go receiver(ch, completed)
     <-completed
	 <-completed
	fmt.Println("Main function completed")

}

func sender(chn chan<- int,completed chan<- bool ) {	
	for i := 0; i <= 5; i++ {
		chn <- i //pushing the data on to the channel
		//time.Sleep(1 * time.Second)
	}
	close(chn)
	completed <- true
}
func receiver(chn <-chan int,completed chan<- bool) {
	for num2 := range chn {
		fmt.Printf("Receiver: %d\n ", num2)
	}
	completed <-true
}

