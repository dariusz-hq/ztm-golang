//--Summary:
//  Create a program that directs vehicles at a mechanic shop
//  to the correct vehicle lift, based on vehicle size.
//
//--Requirements:
//* The shop has lifts for multiple vehicle sizes/types:
//  - Motorcycles: small lifts
//  - Cars: standard lifts
//  - Trucks: large lifts
//* Write a single function to handle all of the vehicles
//  that the shop works on.
//* Vehicles have a model name in addition to the vehicle type:
//  - Example: "Truck" is the vehicle type, "Road Devourer" is a model name
//* Direct at least 1 of each vehicle type to the correct
//  lift, and print out the vehicle information.
//
//--Notes:
//* Use any names for vehicle models

package main

import "fmt"

type Lifter interface {
	Lift()
}

type Motorcycle string
type Car string
type Truck string

func (m Motorcycle) Lift() {
	fmt.Println("Using small lift")
}

func (c Car) Lift() {
	fmt.Println("Using standard lift")
}

func (t Truck) Lift() {
	fmt.Println("Using large lift")
}

func LiftVehicle(l Lifter) {
	fmt.Printf("Lifting vehicle %v \n", l)
}

func main() {
	vehicles := []Lifter{Motorcycle("Ducati"), Car("VW Passat"), Truck("Ford Ranger")}
	for _, v := range vehicles {
		fmt.Printf("Handling lifting for %v \n", v)
		v.Lift()
	}
}
