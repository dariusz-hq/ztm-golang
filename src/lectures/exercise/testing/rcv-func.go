//--Summary:
//  Implement receiver functions to create stat modifications
//  for a video game character.
//
//--Requirements:
//* Implement a player having the following statistics:
//  - Health, Max Health
//  - Energy, Max Energy
//  - Name
//* Implement receiver functions to modify the `Health` and `Energy`
//  statistics of the player.
//  - Print out the statistic change within each function
//  - Execute each function at least once

package main

import "fmt"

type Health int

type Energy int

type Player struct {
	health    Health
	maxHealth Health
	energy    Energy
	maxEnergy Energy
	name      string
}

func (p Player) String() string {
	return fmt.Sprintf("Player %s stats: health %d, maxHealth %d, energy %d, maxEnergy %d", p.name, p.health, p.maxHealth, p.energy, p.maxEnergy)
}

func (p *Player) decrementHealth() {
	if p.health-1 < 0 {
		return
	}
	p.health--
}

func (p *Player) incrementHealth() {
	if p.health+1 > p.maxHealth {
		return
	}
	p.health++
}

func (p *Player) decrementEnergy() {
	if p.energy-1 < 0 {
		return
	}
	p.energy--
}

func (p *Player) incrementEnergy() {
	if p.energy+1 > p.maxEnergy {
		return
	}
	p.energy++
}

func main() {
	p := Player{100, 100, 100, 100, "vader"}
	fmt.Printf("New player %s \n", p.String())

	p.decrementEnergy()
	fmt.Printf("Decremented energy %s \n", p.String())

	p.decrementHealth()
	fmt.Printf("Decremented health %s \n", p.String())

	p.incrementEnergy()
	fmt.Printf("Incremented energy %s \n", p.String())

	p.incrementHealth()
	fmt.Printf("Incremented health %s \n", p.String())
}
