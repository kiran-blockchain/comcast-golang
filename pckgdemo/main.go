package main

import (
	"fmt"

	"github.com/kiran-blockchain/gopackagesdemo/calc"
)

func main(){
	result:=calc.Add(10,20)
	result2:=calc.Mul(10,20)
	fmt.Println(result)
	fmt.Println(result2)
}