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
	
		Carrot.attack()
		Potato.attack()

		fmt.Printf("атака %s: %d\n", Carrot.name,Carrot.health)
		fmt.Printf("атака %s: %d\n\n", Potato.name,Potato.health)

		Carrot.attack()
		Potato.attack()

		fmt.Printf("пополнения здоровье %s: %d\n", Carrot.name,Carrot.health)
		fmt.Printf("пополнения здоровье %s: %d\n\n", Potato.name,Potato.health)

	}
}

func (h *Hero) attack() {
	var attack = rand.Intn(50)
	h.health = h.health - attack
}

func (h *Hero) healthUp() {
	var attack = rand.Intn(50)
 	h.health = h.health + attack
}

