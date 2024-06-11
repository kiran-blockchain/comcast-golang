// package main

// import (
// 	"fmt"
// 	"sync"
// 	"time"
// 	//"time"
// )

// func main() {
// 	var wg sync.WaitGroup
// 	//i :=10;
// 	wg.Add(3)
// 	ch := make(chan int)
// 	go sender(ch, &wg)
// 	go receiver(ch, &wg)
// 	go receiver(ch, &wg)
// 	wg.Wait()
// 	fmt.Println("Main function completed")

// }

// func sender(chn chan<- int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for i := 0; i <= 5; i++ {
// 		chn <- i //pushing the data on to the channel
// 		time.Sleep(1 * time.Second)
// 	}
// 	close(chn)
// }
// func receiver(chn <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num2 := range chn {
// 		fmt.Printf("Receiver: %d\n ", num2)
// 	}
// }
// func receiver2(chn <-chan int, wg *sync.WaitGroup) {
// 	defer wg.Done()
// 	for num2 := range chn {
// 		fmt.Printf("Receiver: %d\n ", num2)
// 	}
// }
