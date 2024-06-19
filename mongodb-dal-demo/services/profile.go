package services

import (
	"context"
	"errors"
	"fmt"
	"mongodemo/entities"
	"mongodemo/interfaces"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileService struct {
	ProfileCollection *mongo.Collection// database table
	Ctx               context.Context // database connection
}

func ProfileServiceInit(profileCollection *mongo.Collection,ctx context.Context) interfaces.IProfile{
	return &ProfileService{ProfileCollection: profileCollection, Ctx:ctx}
}

func (p *ProfileService) CreateProfile(user *entities.Profile) (*entities.Profile, error) {
	user.CreatedAt = time.Now()
	user.UpdatedAt = user.CreatedAt
	user.Email = strings.ToLower(user.Email)
	user.PasswordConfirm = ""
	user.Verified = false
	user.Role = "user"

	// //hashedPassword, _ := utils.HashPassword(user.Password)
	 user.Password = user.Password
	 res, err := p.ProfileCollection.InsertOne(p.Ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	// // Create a unique index for the email field
	// opt := options.Index()
	// opt.SetUnique(true)
	// index := mongo.IndexModel{Keys: bson.M{"email": 1}, Options: opt}

	// if _, err := p.ProfileCollection.Indexes().CreateOne(p.ctx, index); err != nil {
	// 	return nil, errors.New("could not create index for email")
	// }

	 var newUser *entities.Profile
	 query := bson.M{"_id": res.InsertedID}

	err = p.ProfileCollection.FindOne(p.Ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	fmt.Println(newUser.Email)
	return newUser, nil
	//return nil,nil
}

func (p *ProfileService) SearchProfiles()([]*entities.Profile,error){
	var profiles []entities.Profile
	cursor ,err := p.ProfileCollection.Find(p.Ctx,bson.M{"name":"anand"})

	if err != nil {
		return nil, err
	}
	err2:=cursor.All(p.Ctx,&profiles)
	if(err2!=nil){
		fmt.Println("Error iin reading the cursor")
	}
	for _,users := range profiles{
		fmt.Printf("Name: %s \n",users.Name)
	}
	return nil,nil
}