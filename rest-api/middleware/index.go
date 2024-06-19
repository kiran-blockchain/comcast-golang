package middleware

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var jwtKey = []byte("ComcastGOSecretKey")

func ValidateDotken()gin.HandlerFunc{
	return func(c *gin.Context){
		tokenString:= c.GetHeader("Authorization")
		if tokenString ==""{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Authorization Header is not present"})
			c.Abort()
			return
		}
		tokenString = tokenString[len("Bearer "):]
		token,err :=jwt.Parse(tokenString,func(token *jwt.Token)(interface{}, error){
			return jwtKey,nil
		})
		if err!=nil || !token.Valid{
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized Token"})
			c.Abort()
			return
		}
		claims,ok := token.Claims.(jwt.MapClaims)
		if(!ok){
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized Token"})
			c.Abort()
			return
		}
		userName,ok :=claims["username"].(string)
		if(!ok){
			c.JSON(http.StatusUnauthorized,gin.H{"error":"Unauthorized Token"})
			c.Abort()
			return
		}
		c.Set("username",userName)
		c.Next()
	}
}