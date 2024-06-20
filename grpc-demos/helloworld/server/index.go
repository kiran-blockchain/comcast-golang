package main

import (
	"context"
	pb "helloworld/demo"
	"log"
	"net"

	"google.golang.org/grpc"
)



type server struct{
	pb.UnimplementedGreeterServer
}
func(s * server) SayHello(ctx context.Context, input *pb.HelloRequest) (*pb.HelloResponse,error){
	return &pb.HelloResponse{Reply: "Welcome to GRPC " + input.Name},nil
}
//spin up the server
func main(){
	lis,err := net.Listen("tcp",":5001")
	if err!=nil{
		log.Fatalf("failed to listen %v",err)
	}
	s := grpc.NewServer()
	//map the proto buf server on to grpc
	pb.RegisterGreeterServer(s,&server{})
	err = s.Serve(lis)
	 if(err!=nil){
		log.Fatalf("Error in creating server %v",err)
	 }

}