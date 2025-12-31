package main

import (
	"fmt"
	"testing"
)

//--Summary:
//  Copy your rcv-func solution to this directory and write unit tests.
//
//--Requirements:
//* Write unit tests that ensure:
//  - Health & energy can not go above their maximums
//  - Health & energy can not go below 0
//* If any of your  tests fail, make the necessary corrections
//  in the copy of your rcv-func solution file.
//
//--Notes:
//* Use `go test -v ./exercise/testing` to run these specific tests
//

func TestHealthEnergyCannotGoAboveMaximum(t *testing.T) {
	testCases := []struct {
		health    Health
		maxHealth Health
		energy    Energy
		maxEnergy Energy
		name      string
	}{
		{1, 1, 1, 1, "name1"},
		{1, 10, 1, 10, "name1"},
	}

	for _, data := range testCases {
		p := Player{data.health, data.maxHealth, data.energy, data.maxEnergy, data.name}
		for i := Health(0); i <= data.maxHealth+1; i++ {
			p.incrementHealth()
			if p.health > p.maxHealth {
				t.Errorf("Health is %v, want %v", p.health, p.maxHealth)
			}
			fmt.Println(p)
		}
		for i := Energy(0); i <= data.maxEnergy+1; i++ {
			p.incrementEnergy()
			if p.energy > p.maxEnergy {
				t.Errorf("Energy is %v, want %v", p.energy, p.maxEnergy)
			}
			fmt.Println(p)
		}
	}
}

func TestHealthEnergyCannotGoBelowZero(t *testing.T) {
	testCases := []struct {
		health    Health
		maxHealth Health
		energy    Energy
		maxEnergy Energy
		name      string
	}{
		{1, 1, 1, 1, "name1"},
		{1, 10, 1, 10, "name1"},
	}

	for _, data := range testCases {
		p := Player{data.health, data.maxHealth, data.energy, data.maxEnergy, data.name}
		for i := p.maxHealth; i < 0; i-- {
			p.decrementHealth()
			if p.health < 0 {
				t.Errorf("Health is %v, want %v", p.health, 0)
			}
		}
		for i := p.maxEnergy; i < 0; i-- {
			p.decrementEnergy()
			if p.energy < 0 {
				t.Errorf("Energy is %v, want %v", p.energy, 0)
			}
		}
	}
}
