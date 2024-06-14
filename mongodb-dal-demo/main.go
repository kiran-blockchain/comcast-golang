package main

import (
	"context"
	"fmt"
	"log"
	"mongodemo/config"
)

func main(){
	 ctx := context.Background()
	mongoclient,err:=config.ConnectDataBase()
	if(err!=nil){
		log.Fatalf("Error in connecting to the datbase %v",err)
	}
	defer   mongoclient.Disconnect(ctx)
	fmt.Println("Successfully Connected to db....")
}