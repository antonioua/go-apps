package main

import "fmt"

type ninjaStar struct {
	owner string
}

// Define method and implement ninjaWeapon interface
func (n *ninjaStar) attack() {
	fmt.Println(n.owner, " attacks with a ninja star!")
}

type ninjasSword struct {
	owner string
}

// Define method and implement ninjaWeapon interface
func (n *ninjasSword) attack() {
	fmt.Println(n.owner, " attacks with a ninja sword!")
}

// Define interface with method attack for different weapon types
type ninjaWeapon interface {
	attack()
}

func attack(weapon ninjaWeapon) {
	weapon.attack()
}

func main() {
	weapons := []ninjaWeapon{
		&ninjaStar{"Wallace"},
		&ninjasSword{owner: "Wallace"},
		&ninjaStar{"Gromit"},
		&ninjasSword{owner: "Gromit"},
	}

	for _, weapon := range weapons {
		attack(weapon)
		//weapon.attack()
	}
}
