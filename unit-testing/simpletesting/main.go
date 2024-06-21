package main

import (
	//"fmt"
	"strconv"
)

func Verify(input int) string{
	isDivisable := input%3 ==0
	if isDivisable{
		return "true"
	}
	return strconv.Itoa(input)
}
// func main(){
// 	result := Verify(10)
// 	fmt.Println(result)
// }
