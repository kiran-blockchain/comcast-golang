package main

import (
	"context"
	"fmt"
	"log"
	"mongodemo/config"
	"mongodemo/constants"
	"mongodemo/entities"
	"mongodemo/services"
)

func main(){
	 ctx := context.Background()
	 mongoClient,err :=config.ConnectDataBase()
	 profileCollection:= mongoClient.Database(constants.DatabaseName).Collection("profiles")
	 profileService := services.ProfileServiceInit(profileCollection, ctx)
	 profile:=entities.Profile{
		Name: "kiran",
		Email: "Kiran@gmail.com",
	 }
	 result,err:= profileService.CreateProfile(&profile)
	 if(result!=nil){
		fmt.Println("Profile Created")
	 }
	 
	if(err!=nil){
		log.Fatalf("Error in connecting to the datbase %v",err)
	}
	defer   mongoClient.Disconnect(ctx)
	fmt.Println("Successfully Connected to db....")
}

