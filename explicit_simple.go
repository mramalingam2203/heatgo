package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math"
	"os"
)

var pi float64 = 4 * math.Atan(1.0)

func explicit_matlab(ntime int, nx int, dt float64, dx float64, T_left float64, T_right float64, r float64) {
	t := make([]float64, ntime)
	T := make([]float64, nx)
	T0 := make([]float64, nx)

	// Open the file for writing
	file, err := os.Create("output.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	for i := 0; i < ntime; i++ {
		t[i] = float64(i) * dt
	}

	for j := 0; j < nx; j++ {
		T0[j] = T_left
	}

	for i := 0; i < ntime; i++ {

		for j := 0; j < nx; j++ {
			T[j] = T0[j]
		}

		for k := 0; k < nx; k++ {
			if k == 0 {
				T[k] = T_left
			} else if k == int(nx-1) {
				T[k] = T_right
			} else {
				T[k] = T[k] + r*(T[k+1]-2*T[k]+T[k-1])
			}

		}
		k := 0
		// Convert each float in the slice to a string and write a single row
		for _, value := range T {
			row := []string{fmt.Sprintf("%.2f %.2f %.3f", float64(i)*dt, float64(k)*dx, value)} // Format the float value to two decimal places
			err := writer.Write(row)
			if err != nil {
				log.Fatal(err)
			}
			k++
		}

		// Flush any buffered data to the underlying writer (file)
		writer.Flush()

		// Check for any errors that occurred during the flush
		if err := writer.Error(); err != nil {
			log.Fatal(err)
		}

		for j := 0; j < nx; j++ {
			fmt.Sprintf("%[2]d %[1]d\n", j, T[j])
		}

		for j := 0; j < nx; j++ {
			T0[j] = T[j]
		}

	}

}
