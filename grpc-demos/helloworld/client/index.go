package main

import (
	"context"
	"fmt"
	pb "helloworld/demo"
	"log"
	"time"

	"google.golang.org/grpc"
	
)
func main(){
  //create a connection to the server
  conn, err := grpc.NewClient(":5001",grpc.WithInsecure())
   if err!=nil{
	 log.Fatalf("Error in conencting to server %v",err)
   }
   defer conn.Close()
   c := pb.NewGreeterClient(conn)
   ctx, cancel := context.WithTimeout(context.Background(),5*time.Second)
   defer cancel()
   reply, err2 := c.SayHello(ctx,&pb.HelloRequest{Name: "Kiran"})
   if err2!=nil{
	log.Fatalf("Error in sending response %v", err2)
   }
   fmt.Println(reply)
}