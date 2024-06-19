package controllers

import (
	"net/http"
	"rest-api/interfaces"
	"rest-api/models"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService interfaces.IUser
}

func UserControllerInit(userService interfaces.IUser) UserController {
	return UserController{userService}
}

func (u *UserController) Login(c *gin.Context){
	var user models.User
	//convert the posted object from the UI to JSON.
	 err:=c.BindJSON(&user)
	 if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	 }
	 result,err2:= u.UserService.Login(&user)
	 if(err2!=nil){
		c.JSON(http.StatusUnauthorized,gin.H{"error":err2.Error()})
		return
	 }
	 c.JSON(http.StatusOK,gin.H{"message":"Login Successful","token":result})
	
}

func (u *UserController) Register(c *gin.Context){
	var user models.User
	//convert the posted object from the UI to JSON.
	 err:=c.BindJSON(&user)
	 if err!=nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	 }
	 err2:= u.UserService.Register(&user)
	 if(err2!=nil){
		c.JSON(http.StatusUnauthorized,gin.H{"error":err2.Error()})
		return
	 }
	 c.JSON(http.StatusOK,gin.H{"message":"User Created Successfully"})
	
}

func (u *UserController) FetchProfile(c *gin.Context) {


}