package main

import (
	"context"

	pb "grpc-demo2/protobuf"
	"log"
	"time"

	"google.golang.org/grpc"
)
func main(){
  //create a connection to the server
  conn, err := grpc.Dial(":5001",grpc.WithInsecure())
   if err!=nil{
	 log.Fatalf("Error in conencting to server %v",err)
   }
   defer conn.Close()
   c := pb.NewChatClient(conn)
   
   ctx, cancel := context.WithTimeout(context.Background(),2*time.Second)
   defer cancel()
   stream,err  := c.JoinChat(ctx)
   if err!=nil{
	log.Fatalf("Error in conencting to server %v",err)
  }
   go func(){

	for _,msg := range[]pb.ChatMessage{
		{Username: "Client1", Message: "hello"},
		{Username: "Client2", Message: "How are you"},
		{Username: "Client3", Message: "Good Bye"},
	}{
		if err:= stream.Send(&msg);err!=nil{
			log.Fatalf("Error in reciving data  %v",err)
		}
		time.Sleep(3* time.Second)
	}
   }()
   for {
	resp, err := stream.Recv()
	if err!=nil{
		log.Fatalf("Error in reciving to server %v",err)
		break
	  }
	  log.Printf("Recived data from server : %s", resp.Message)
   }
  
}