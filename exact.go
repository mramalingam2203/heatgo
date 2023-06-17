package main

import (
	"fmt"
	"math"
)

func exact(nodes int, Diff float64, output_time float64, delta_x float64, length float64, T_sur float64,
	T_init float64) {
	var series float64
	var acc int32 = 0
	exact := make([]float64, nodes)

	for i := 0; i < nodes; i++ {
		series = 0
		for m := 1; m <= int(acc); m++ {
			series += math.Exp(-Diff*output_time*math.Pow((float64(m)*math.Pi)/length, 2)) * ((1.0 - math.Pow(-1.0, float64(m))) / (float64(m) * math.Pi)) * math.Sin((float64(m)*math.Pi*float64(i)*delta_x)/length)
		}
	}

	for i := 0; i < nodes; i++ {
		exact[i] = T_sur + 2*(T_init-T_sur)*series
		fmt.Println(exact[i])
	}

}
