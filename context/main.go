package main

import (
	"context"
	"fmt"
	"time"
)
func doWork(ctx context.Context){
	for {
		select{
		case <-time.After(1*time.Second):
			fmt.Println("Workingg.....")
		case <-ctx.Done():
			fmt.Println("Context Cancelled")
			return
		}
	}
}

func main (){
	ctx,cancel := context.WithDeadline(context.Background(),time.Now().Add(5*time.Second))
	defer cancel()
	go doWork(ctx)
	time.Sleep(10*time.Second)
}