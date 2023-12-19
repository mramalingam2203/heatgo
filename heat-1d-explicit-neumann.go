package main

import (
	"fmt"
	"math"
)

func initialCondition(x float64) float64 {
	// Initial condition: u(x, 0)
	// You can change this function based on your specific initial condition
	return math.Sin(math.Pi * x)
}

func neumannBC(u []float64, dx float64) {
	// Neumann boundary conditions: du/dx = 0 at x=0 and x=L
	u[0] = u[1]
	u[len(u)-1] = u[len(u)-2]
}

func main() {
	// Parameters
	NX := 50      // Number of spatial points
	NT := 100     // Number of time steps
	LX := 1.0     // Length of the spatial domain
	LT := 0.1     // Total simulation time
	alpha := 0.01 // Thermal diffusivity

	dx := LX / float64(NX-1) // Spatial step size
	dt := LT / float64(NT)   // Time step size

	// Solution array
	u := make([]float64, NX)

	// Initialize the solution array with the initial condition
	for i := 0; i < NX; i++ {
		u[i] = initialCondition(float64(i) * dx)
	}

	// Time-stepping loop
	for n := 0; n < NT; n++ {
		// Create a copy of the solution array for the current time step
		uNew := make([]float64, len(u))
		copy(uNew, u)

		// Update the solution using explicit finite differences
		for i := 1; i < len(u)-1; i++ {
			uNew[i] = u[i] + alpha*dt/(dx*dx)*(u[i-1]-2*u[i]+u[i+1])
		}

		// Apply Neumann boundary conditions
		neumannBC(uNew, dx)

		// Update the solution array for the next time step
		copy(u, uNew)
	}

	// Print the final solution or perform further analysis as needed
	for i := 0; i < NX; i++ {
		fmt.Printf("%f ", u[i])
	}
}

