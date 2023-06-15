package main

import (
	"fmt"
	"math"
)

var pi float64 = 4 * math.Atan(1.0)

func explicit_matlab(ntime int32, nx int32, dt float32, T_left float32, T_right float32, r float32) {
	t := make([]float32, ntime)
	T := make([]float32, nx)
	T0 := make([]float32, nx)

	for i := 0; i < int(ntime); i++ {
		t[i] = float32(i) * dt
	}

	for j := 0; j < int(nx); j++ {
		T0[j] = T_right
	}

	for i := 0; i < int(ntime); i++ {

		for j := 0; j < int(nx); j++ {
			T[j] = T0[j]
		}

		for k := 0; k < int(nx); k++ {
			if k == 0 {
				T[k] = T_left
			} else if k == int(nx-1) {
				T[k] = T_right
			} else {
				T[k] = T[k] + r*(T[k+1]-2*T[k]+T[k-1])
			}

		}

		for j := 0; j < int(nx); j++ {
			T0[j] = T0[j]
		}

	}

	// fmt.Println(T)
}

func exact(nodes int32, Diff float64, output_time float64, delta_x float64, length float64, T_sur float64,
	T_init float64) {
	var series float64
	var acc int32 = 0
	exact := make([]float64, nodes)

	for i := 0; i < int(nodes); i++ {
		series = 0
		for m := 1; m <= int(acc); m++ {
			series += math.Exp(-Diff*output_time*math.Pow((float64(m)*math.Pi)/length, 2)) * ((1.0 - math.Pow(-1.0, float64(m))) / (float64(m) * math.Pi)) * math.Sin((float64(m)*math.Pi*float64(i)*delta_x)/length)
		}
	}

	for i := 0; i < int(nodes); i++ {
		exact[i] = T_sur + 2*(T_init-T_sur)*series
		fmt.Println(exact[i])
	}

}

func DuFort_Frankel_Explicit_Scheme(delta_x float64, delta_t float64, output_time float64, nodes int32, r float64, T_sur float64, T_init float64) {

	numerical := make([][]float64, 0.0)

	for n := 0; n <= int(output_time/delta_t); n++ {
		numerical[0][n] = T_sur       // BC at 0 and all n (node #0)
		numerical[nodes-1][n] = T_sur // BC at 31 and all n (node #last_node)
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #1
		numerical[i][0] = T_init // IC at n = 0, for all i, except node node i = 0 and i = nodes
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #2
		numerical[i][1] = r*numerical[i-1][0] + (1.0-(2.0*r))*numerical[i][0] + r*numerical[i+1][0] // IC at n = 1, for all i, except node node i = 0 and i = nodes
	}

	for n := 1; n < int(output_time/delta_t); n++ {
		for i := 1; i < int(nodes-1); i++ {
			numerical[i][n+1] = ((2.0*r)/(1.0+2.0*r))*numerical[i-1][n] + ((2.0*r)/(1.0+2.0*r))*numerical[i+1][n] + ((1.0-2.0*r)/(1.0+2.0*r))*numerical[i][n-1]
		}
	}
}

func Richardson_Explicit_Scheme(delta_x float64, delta_t float64, output_time float64, nodes int32, r float64, T_sur float64, T_init float64) {

	numerical := make([][]float64, 0.0)

	for n := 0; n <= int(output_time/delta_t); n++ {
		numerical[0][n] = T_sur       // BC at 0 and all n (node #0)
		numerical[nodes-1][n] = T_sur // BC at 31 and all n (node #last_node)
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #1
		numerical[i][0] = T_init // IC at n = 0, for all i, except node node i = 0 and i = nodes
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #2
		numerical[i][1] = r*numerical[i-1][0] + (1.0-(2.0*r))*numerical[i][0] + r*numerical[i+1][0] // IC at n = 1, for all i, except node node i = 0 and i = nodes
	}

	for n := 1; n < int(output_time/delta_t); n++ {
		for i := 1; i < int(nodes-1); i++ {
			numerical[i][n+1] = 2*r*numerical[i-1][n] - 4*r*numerical[i][n] + 2*r*numerical[i+1][n] + numerical[i][n-1]

		}
	}

}

func Laasonen_Simple_Implicit_Scheme(numerical *[][] float64, delta_x float64, delta_t float64, double output_time float64, nodes int, r float64, T_sur float64, T_init float64) {

  
	for n := 0; n <= int(output_time / delta_t); n++ {
		numerical[0][n] = T_sur;
		numerical[nodes - 1][n] = T_sur;
	}

	for i := 1; i < (nodes - 1); i++ {
		numerical[i][0] = T_init;					
	}

}

func Crank_Nicholson_Implicit_Scheme(numerical *[][] float64, delta_x float64, delta_t float64, double output_time float64, nodes int, r float64, T_sur float64, T_init float64) {

}


func TDMA_Solver(lower_diag []float64, main_diag []float64, upper_diag []float64, b []float64, nodes int) {
	nodes = nodes - 2
	nodes--
	upper_diag[0] /= main_diag[0]
	b[0] /= main_diag[0]

	for i := 1; i < nodes; i++ {
		upper_diag[i] /= main_diag[i] - lower_diag[i]*upper_diag[i-1]
		b[i] = (b[i] - lower_diag[i]*b[i-1]) / (main_diag[i] - lower_diag[i]*upper_diag[i-1])
	}

	b[nodes] = (b[nodes] - lower_diag[nodes]*b[nodes-1]) / (main_diag[nodes] - lower_diag[nodes]*upper_diag[nodes-1])

	for i := nodes; i > 0; i-- {
		b[i] -= upper_diag[i] * b[i+1]
	}

}

func main() {
	var xmin, xmax, dX, dt, maxtime, rvalue, alpha, Tl, Tr float32
	var nX, nt int32
	xmin = 0.0
	xmax = 1.0
	nX = 100
	maxtime = 25
	nt = 100
	alpha = 1e-4
	Tl = 0.0
	Tr = 100.0

	dX = float32(xmax-xmin) / float32(nX-1)
	dt = float32(maxtime) / float32(nt)
	fmt.Println(dX, dt)

	rvalue = alpha * dt / dX * dX
	fmt.Println(rvalue)

	explicit_matlab(nt, nX, dt, Tl, Tr, rvalue)

}
