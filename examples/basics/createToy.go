package main

import "fmt"

// Toy represents a generic toy with name and price.
type Toy struct {
	Name  string
	Price float64
}

// NewToy creates a new Toy instance with the given name and price.
func NewToy(name string, price float64) *Toy {
	return &Toy{
		Name:  name,
		Price: price,
	}
}

// Play simulates playing with the toy.
func (t *Toy) Play() {
	fmt.Printf("Playing with the %s toy!\n", t.Name)
}

// SetPrice sets the price of the toy.
func (t *Toy) SetPrice(price float64) {
	t.Price = price
}

// GetPrice returns the price of the toy.
func (t *Toy) GetPrice() float64 {
	return t.Price
}

// ToyCar represents a toy car that inherits from the Toy struct.
type ToyCar struct {
	Toy    // Embedded Toy struct
	Model  string
	Engine string
}

// NewToyCar creates a new ToyCar instance with the given name, price, model, and engine.
func NewToyCar(name string, price float64, model string, engine string) *ToyCar {
	return &ToyCar{
		Toy:    *NewToy(name, price), // Embed Toy struct
		Model:  model,
		Engine: engine,
	}
}

// Drive simulates driving the toy car.
func (tc *ToyCar) Drive() {
	fmt.Printf("Driving the %s toy car with model %s and engine %s!\n", tc.Name, tc.Model, tc.Engine)
}

func main() {
	// Create a new ToyCar instance
	toyCar := NewToyCar("Toy Car", 19.99, "Sports", "Electric")

	// Play with the toy car (inherited from Toy struct)
	toyCar.Play()

	// Get the price of the toy car (inherited from Toy struct)
	fmt.Printf("The price of the %s toy car is $%.2f\n", toyCar.Name, toyCar.GetPrice())

	// Set a new price for the toy car (inherited from Toy struct)
	toyCar.SetPrice(24.99)

	// Get the updated price of the toy car (inherited from Toy struct)
	fmt.Printf("The updated price of the %s toy car is $%.2f\n", toyCar.Name, toyCar.GetPrice())

	// Drive the toy car (specific to ToyCar struct)
	toyCar.Drive()
}
