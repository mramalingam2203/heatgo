package main

import (
	"fmt"
	"math"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/lapack/lapack64"
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

	// Construct the coefficient matrix A for the implicit scheme
	A := constructMatrix(NX, alpha, dt, dx)

	// Solution array
	u := make([]float64, NX)

	// Initialize the solution array with the initial condition
	for i := 0; i < NX; i++ {
		u[i] = initialCondition(float64(i) * dx)
	}

	// Time-stepping loop
	for n := 0; n < NT; n++ {
		// Construct the right-hand side vector b
		b := constructRHS(u, alpha, dt, dx)

		// Solve the linear system Au = b for the new solution u
		solveLinearSystem(A, b, u)
	}

	// Print the final solution or perform further analysis as needed
	for i := 0; i < NX; i++ {
		fmt.Printf("%f ", u[i])
	}
}

// constructMatrix constructs the coefficient matrix A for the implicit scheme
func constructMatrix(NX int, alpha, dt, dx float64) *mat.Dense {
	size := NX - 2
	A := mat.NewDense(size, size, nil)

	// Diagonal elements
	for i := 0; i < size; i++ {
		A.Set(i, i, 1.0+2.0*alpha*dt/(dx*dx))
	}

	// Off-diagonal elements
	for i := 0; i < size-1; i++ {
		A.Set(i, i+1, -alpha*dt/(dx*dx))
		A.Set(i+1, i, -alpha*dt/(dx*dx))
	}

	return A
}

// constructRHS constructs the right-hand side vector b for the implicit scheme
func constructRHS(u []float64, alpha, dt, dx float64) []float64 {
	size := len(u) - 2
	b := make([]float64, size)

	for i := 1; i < len(u)-1; i++ {
		b[i-1] = u[i] + alpha*dt/(dx*dx)*(u[i-1]-2*u[i]+u[i+1])
	}

	return b
}

// solveLinearSystem solves the linear system Au = b using LAPACK
func solveLinearSystem(A *mat.Dense, b []float64, u []float64) {
	// Convert b to a matrix
	bVec := mat.NewVecDense(len(b), b)

	// Solve the linear system using LAPACK
	var pivots lapack64.Pivots
	err := lapack64.LU(A, pivots, bVec)
	if err != nil {
		fmt.Println("Error solving linear system:", err)
	}

	// Copy the solution back to the u array
	copy(u[1:len(u)-1], bVec.RawVector().Data)
}

