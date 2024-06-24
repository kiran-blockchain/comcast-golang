package main

import (
	"context"
	"fmt"
	"log"
	"rest-api/config"
	"rest-api/constants"
	"rest-api/controllers"
	"rest-api/routes"
	"rest-api/services"

	//	"rest-api/services"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)


var (
	mongoclient *mongo.Client
	ctx         context.Context
	server         *gin.Engine
)
func initRoutes(){
  routes.Default(server)
  
}

func initApp(mongoClient *mongo.Client){
	ctx = context.TODO()

	//profile service is dependent on table and the database .
	profileCollection := mongoClient.Database(constants.DatabaseName).Collection("profiles")
	userCollection := mongoClient.Database(constants.DatabaseName).Collection("users")
	
	//pass the table and database to the profile service as parameters
	profileService := services.NewProfileServiceInit(profileCollection, ctx)
	userService := services.UserServiceInit(userCollection,ctx)
	//inject the profile service as a depdency to the controller.
	profileController := controllers.InitProfileController(profileService)
	userController := controllers.UserControllerInit(userService)
	routes.ProfileRoute(server,profileController)
	routes.UserRoutes(server,userController)
}


func main(){
	server = gin.Default()
	
	pprof.Register(server, "dev/pprof")
	mongoclient,err :=config.ConnectDataBase()
	defer   mongoclient.Disconnect(ctx)
	if err!=nil{
		panic(err)
	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port",constants.Port)
	log.Fatal(server.Run(constants.Port))
}