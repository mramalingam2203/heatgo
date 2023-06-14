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
	for i := 0; i < int(nodes); i++ {
		series = 0
		for m := 1; m <= int(acc); m++ {
			series += math.Exp(-Diff*output_time*math.Pow((float64(m)*math.Pi)/length, 2)) * ((1.0 - math.Pow(-1.0, float64(m))) / (float64(m) * math.Pi)) * math.Sin((float64(m)*math.Pi*float64(i)*delta_x)/length)
		}
	}

}

func explicit() {

}

func implicit() {

}
func cranknic() {

}
func ADI() {

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
