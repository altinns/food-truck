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

var foodOrder string
var foodAmount int 
var drinkOrder string
 var drinkAmount int 

func  FdOrder(x string,a int)int{
 amount:=a

return Foodlist[x]-amount
	
}





func main(){


//var drinkOrder int
 //var drinkAmount int 
 
fmt.Printf("Welcome to :Original Taste")
fmt.Printf("Order menu: \nFoods: 1.Burger 2.Toast \n")
fmt.Scan(&foodOrder)
fmt.Printf("How many? \n")
fmt.Scan(&foodAmount)
fmt.Println("Drinks: 1.Cola 2.Fanta")
fmt.Scan(&drinkOrder)
fmt.Printf("How many?")
fmt.Scan(&drinkAmount)
fmt.Println(Foodlist)

food:=FdOrder(foodOrder,foodAmount)
Foodlist[foodOrder]=food
// fmt.Println(food)
fmt.Printf( "Foodlist[%s]=%d",foodOrder,Foodlist[foodOrder]) 


drink:=FdOrder(drinkOrder,drinkAmount)
Foodlist[drinkOrder]=drink
//fmt.Println(drink)
fmt.Printf( "Foodlist[%s]=%d",drinkOrder,Foodlist)



fmt.Printf(Foodlist)



//order,ok:=Foodlist[foodOrder]
//if ok{
//return order
//}
//Foodlist["burger"]=Foodlist["burger"]-foodAmount //this works
//fmt.Println(Foodlist["burger"])
}