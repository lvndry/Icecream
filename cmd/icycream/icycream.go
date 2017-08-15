package main

import (
	"fmt"
	"log"

	gini "github.com/lvndry/IcyCream/pkg/supply/supply.go"
)

var (
	flavors = make(map[string]int)
	extras  = make(map[string]bool)
)

func main() {
	var numofballs int
	log.Println("Welcome in IceCreamLand !")
	log.Println("Filling the flavors..")
	getFlavors()
	log.Println(flavors)
	log.Println("Filling the extras..")
	getExtras()
	log.Println(extras)
	log.Println("How many balls do you want : ")
	fmt.Scanf("%d", &numofballs)
	log.Printf("You want %d balls", numofballs)
	gini.FillFlavors()
	log.Println(flavors)
}

func getFlavors() {
	flavors["bastani"]
	flavors["bacon"]
	flavors["chocolate"]
	flavors["coconut"]
	flavors["pistacho"]
	flavors["vanilla"]
}

func getExtras() {
	extras["caramel"] = false
	extras["strawberry"] = false
	extras["chantilly"] = false
}
