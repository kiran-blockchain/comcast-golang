package main

import (
	"goormdemo/entities"
	"log"
)


func main(){
	connStr := "host=SG-storm-bath-6462-5504-pgsql-master.servers.mongodirector.com user=ecom2 password=Ecom2#2024 dbname=ecom sslmode=disable"
	profileService:= entities.NewProfileService(connStr)
	//newProfile := &entities.Profile{Username: "Kiran PVS",Email: "kiran@gmail.com"}
	//err:= profileService.Create(newProfile)
	//err:=profileService.Update(1,"kiran2@gmail.com")
	err:= profileService.GetOrdersByUserId()
	if(err!=nil){
		log.Fatalf("Error in creating the user %v",err)
	}
	
}