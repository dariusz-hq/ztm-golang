//--Summary:
//  Create a program that can activate and deactivate security tags
//  on products.
//
//--Requirements:

//* Create functions to activate and deactivate security tags using pointers
//* Create a checkout() function which can deactivate all tags in a slice
//* Perform the following:
//  - Create at least 4 items, all with active security tags
//  - Store them in a slice or array
//  - Deactivate any one security tag in the array/slice
//  - Call the checkout() function to deactivate all tags
//  - Print out the array/slice after each change

package main

import "fmt"

type Product struct {
	name string
	tag  bool
}

func activate(item *Product) {
	item.tag = true
}

func deactivate(item *Product) {
	item.tag = false
}

func checkout(products []Product) {
	for i := range products {
		deactivate(&products[i])
	}
}

func main() {
	//* Create a structure to store items and their security tag state
	//  - Security tags have two states: active (true) and inactive (false)
	products := []Product{
		{"item1", true},
		{"item2", true},
		{"item3", true},
		{"item4", true},
	}
	fmt.Println("products", products)

	deactivate(&products[0])
	fmt.Println("after deactivating", products)

	checkout(products)
	fmt.Println("after checkout", products)

}
