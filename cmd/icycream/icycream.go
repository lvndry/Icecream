package main

import (
	"log"

	"github.com/lvndry/Icycream/pkg/supply"
)

var (
	flavors = make(map[string]int)
	extras  = make(map[string]bool)
)

func main() {
	log.Println("Welcome in IceCreamLand !")

	log.Println("Let's buy some icecream flavors..")
	getFlavors()
	//log.Println(flavors)

	log.Println("and some extras too..")
	getExtras()
	//log.Println(extras)

	log.Println("Now let's get them on the fridge!")
	supply.FillFlavors(flavors)
	log.Println("Done")
	//log.Println(flavors)
	supply.FillExtras(extras)
	log.Println("Done")
	//log.Println(extras)

	log.Println("Switching on phone..")
	router := route.Init()
}

func getFlavors() {
	flavors["bastani"] = 0
	flavors["bacon"] = 0
	flavors["chocolate"] = 0
	flavors["coconut"] = 0
	flavors["pistacho"] = 0
	flavors["vanilla"] = 0
	log.Println("Done!")
}

func getExtras() {
	extras["caramel"] = false
	extras["strawberry"] = false
	extras["chantilly"] = false
	log.Println("Done!")
}
