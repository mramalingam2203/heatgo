
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
		// Copy the current solution to a new array
		unew := make([]float64, NX)
		copy(unew, u)

		for i := 1; i < NX-1; i++ {
			// Explicit finite difference scheme (1D heat equation)
			unew[i] = u[i] + alpha*dt/(dx*dx)*(u[i+1]-2*u[i]+u[i-1])
		}

		// Update the solution array
		copy(u, unew)
	}

	// Print the final solution or perform further analysis as needed
	for i := 0; i < NX; i++ {
		fmt.Printf("%f ", u[i])
	}
}
