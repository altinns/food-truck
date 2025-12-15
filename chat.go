package main

import (
	"fmt"
)

var Foodlist=map[string]int{

	"burger":100,
"toast":100,
"Cola":100,
"Fanta":100,
}



func  FdOrder(x string,a int)int{
 amount:=a
result:=Foodlist[x]

	return result-amount
}





func main(){

var foodOrder string
var foodAmount int 
//var drinkOrder int
 //var drinkAmount int 
 
fmt.Printf("Welcome to :Original Taste")
fmt.Printf("Order menu: \n Foods: 1.Burger 2.Toast \n")
fmt.Scan(&foodOrder)
fmt.Printf("How many?")
fmt.Scan(&foodAmount)
//fmt.Println("Drinks: 1.Cola 2.Fanta")
//fmt.Scan(&drinkOrder)
//fmt.Printf("How many?")
//fmt.Scan(&drinkAmount)
//fmt.Println(Foodlist)

result:=FdOrder(foodOrder,foodAmount)

fmt.Println(result)



}