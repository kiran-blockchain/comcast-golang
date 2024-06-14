package main

import (
	"context"
	"fmt"
	"log"
	"mongodemo/config"
	"mongodemo/constants"
	"mongodemo/entities"
	"mongodemo/interfaces"
	"mongodemo/services"

	"go.mongodb.org/mongo-driver/mongo"
)
var(
	ctx  context.Context
	mongoClient *mongo.Client
	profileCollection *mongo.Collection
	productCollection *mongo.Collection
	profileService interfaces.IProfile
	err error
)

func init(){

	ctx = context.Background()
	mongoClient,err =config.ConnectDataBase()
	if(err!=nil){
		log.Fatalf("Error in connecting to the datbase %v",err)
	}
	profileCollection = mongoClient.Database(constants.DatabaseName).Collection("profiles")
	profileService = services.ProfileServiceInit(profileCollection, ctx)
}

func main(){
	 
	 profile:=entities.Profile{
		Name: "kiran",
		Email: "KiranTest3@gmail.com",
	 }
	 result,err2:= profileService.CreateProfile(&profile)
	 if(result!=nil){
		fmt.Println("Profile Created")
	 }
	 if(err2!=nil){
		fmt.Println("error in creating profile")
	 }
	 
	
	defer   mongoClient.Disconnect(ctx)
	fmt.Println("Successfully Connected to db....")
}

