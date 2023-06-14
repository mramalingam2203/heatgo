package main

import "fmt"

func explicit_matlab() {

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
	var xmin, xmax, dx, dt, maxtime, r, alpha float32
	var nx, ntime int32
	xmin = 0.0
	xmax = 1.0
	nx = 100
	maxtime = 25
	ntime = 100
	alpha = 1e-4

	dx = float32(xmax-xmin) / float32(nx-1)
	dt = float32(maxtime) / float32(ntime)
	fmt.Println(dx, dt)

	r = alpha * dt / dx * dx
	fmt.Println(r)

}
