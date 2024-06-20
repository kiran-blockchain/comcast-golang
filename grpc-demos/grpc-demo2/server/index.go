package main

import (
	pb "grpc-demo2/protobuf"
	"log"
	"net"

	"google.golang.org/grpc"
)	



type server struct{
	pb.UnimplementedChatServer
}

func (s *server) JoinChat(stream pb.Chat_JoinChatServer) error{
	for{
		msg, err := stream.Recv()
		 if err!=nil {
			log.Printf("error in receving the messages %v", err)
		 }
		 log.Printf("Received Message from %s: %s",msg.Username,msg.Message)

		 if err := stream.Send(&pb.ChatMessage{
			Username: "Server",
			Message: "Received: "+msg.Message,
		 });err !=nil{
			log.Printf("Error Sending Message: %v",err)
		 }

	}
	return nil
}
func main(){
	lis,err := net.Listen("tcp",":5001")
	if err!=nil{
		log.Fatalf("failed to listen %v",err)
	}
	s := grpc.NewServer()
	
	//map the proto buf server on to grpc
	pb.RegisterChatServer(s,&server{})
	err = s.Serve(lis)
	 if(err!=nil){
		log.Fatalf("Error in creating server %v",err)
	 }

}
