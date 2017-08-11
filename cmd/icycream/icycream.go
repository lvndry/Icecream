package main

import (
	"fmt"
	"log"
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
}

func getFlavors() {
	flavors["bastani"] = 10
	flavors["bacon"] = 10
	flavors["chocolate"] = 10
	flavors["coconut"] = 10
	flavors["pistacho"] = 10
	flavors["vanilla"] = 10
}

func getExtras() {
	extras["caramel"] = false
	extras["strawberry"] = false
	extras["chantilly"] = false
}
