package services

import (
	"context"
	"errors"
	"fmt"
	"rest-api/interfaces"
	"rest-api/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserCollection *mongo.Collection // database table
	ctx            context.Context   // database connection
}

func UserServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.IUser {

	return &UserService{collection, ctx}
}
func (u *UserService) Register(user *models.User) error {
	//check if the user is existing?
	existingUser := models.User{}
	filter :=bson.M{"username":user.Username}
	err:= u.UserCollection.FindOne(u.ctx,filter).Decode(&existingUser)
	if err==nil{
		fmt.Println("user already exists")
		return errors.New("user already exists")
	}
	//if the user is not existing create the user
	//hash the password
	hashedPwd, errInEncryption:= bcrypt.GenerateFromPassword([]byte(user.Password),bcrypt.DefaultCost) 
	 if(errInEncryption!=nil){
		fmt.Println("Erorr in incryption")
		return errors.New("Error in Ecrypting Password")
	 }
	 user.Password= string(hashedPwd)
	_,err = u.UserCollection.InsertOne(u.ctx,user)
	
	if(err!=nil){
		return err
	}
	return nil
}
func (u *UserService) Login(user *models.User) (*models.User, error) {
	 var result *models.User
	filter :=bson.M{"username":user.Username,"password":user.Password}

	
	err:= u.UserCollection.FindOne(u.ctx,filter).Decode(&result)
	err2:= bcrypt.CompareHashAndPassword([]byte(user.Password),[]byte(result.Password))
	
	if err!=nil || err2!=nil{
		return nil, errors.New("invalid Credentials")
	}
	return result, nil
}
func (u *UserService) FetchProfile(id string) (*models.User, error) {
	objId,err :=primitive.ObjectIDFromHex(id)
	if(err!=nil){
		return nil,errors.New("invalid user")
	}
	filter :=bson.M{"_id": objId}
	var existingUser *models.User
	err = u.UserCollection.FindOne(u.ctx,filter).Decode(&existingUser)
	if err!=nil{
		return nil,errors.New("User doesnot exist")
	}
	return existingUser,nil
}
