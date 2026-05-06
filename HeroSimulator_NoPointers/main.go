package main

import (
	"fmt"
	"math/rand"
)

	type Hero struct{
		name string
		health int
	}
	var Potato Hero = Hero{
		name: "Potato",
		health: 100,
	}	

	var Carrot Hero = Hero{
		name: "Carrot",
		health: 100,
	}
func main()  {
	for true{
		fmt.Printf("%s: %d%%\n", Carrot.name,Carrot.health)
		fmt.Printf("%s: %d%%\n\n", Potato.name,Potato.health)

		if Potato.health <= 0  {
			fmt.Printf("WIN: %s", Carrot.name)
			break
		}

		if Carrot.health <= 0  {
			fmt.Printf("WIN: %s", Potato.name)
			break
		}
	
		Carrot.health = attack(Carrot)
		Potato.health =attack(Potato)

		fmt.Printf("атака %s: %d\n", Carrot.name,Carrot.health)
		fmt.Printf("атака %s: %d\n\n", Potato.name,Potato.health)

		Carrot.health = health(Carrot)
		Potato.health = health(Potato)

		fmt.Printf("пополнения здоровье %s: %d\n", Carrot.name,Carrot.health)
		fmt.Printf("пополнения здоровье %s: %d\n\n", Potato.name,Potato.health)

	}
}

func attack(h Hero) int{
	var attack = rand.Intn(50)
	h.health = h.health - attack
	return h.health
}

func health(h Hero) int{
	var attack = rand.Intn(50)
 	h.health = h.health + attack
	return h.health
}
