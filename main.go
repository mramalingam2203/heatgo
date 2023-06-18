/* Finite Difference Solutions for the transient 1D heat equation trying an exhaustive approach of all available techniques

Technique adopted from Tannehill & Pletcher famous CFD book
*/

package main

import (
	"fmt"
	_ "image"
	_ "image/color"
	_ "image/draw"
	_ "image/png"
	_ "log"
	_ "math/rand"
	_ "os"
	_ "sort"
	_ "time"

	_ "github.com/ajstarks/svgo"
	_ "github.com/vdobler/chart"
	_ "github.com/vdobler/chart/imgg"
	_ "github.com/vdobler/chart/svgg"
	_ "github.com/vdobler/chart/txtg"
)

func main() {

	// Variables meant for 1-D calculations only

	/*
		xmin - minimum value of x
		xmax - maximum value of x
		dX - distance inrements
		dt - time increment
		maxtime - total time extent
		r - Courant Number = alpha*dt/dx*dx
		alpha - thermal diffusivity = rho*cp/K
		Tl - Non-dimensional Temperature at left end
		Tr - non-dimensional temperature at right end
		nx - total number of nodes in X-direction
		nt - total number of time steps
	*/
	var xmin, xmax, dX, dt, maxtime, rvalue, alpha, Tl, Tr float64
	var nX, nt int

	alpha = 0.05
	xmin = 0
	xmax = 1.0
	nX = 5
	Tl = 0.0
	Tr = 300.0
	dX = (xmax - xmin) / float64(nX-1)
	dt = 1
	maxtime = 10
	nt = int(maxtime / dt)

	dX = float64(xmax-xmin) / float64(nX-1)
	dt = float64(maxtime) / float64(nt)
	fmt.Println(dX, dt)

	rvalue = alpha * dt / dX * dX
	fmt.Println(rvalue)

	explicit_matlab(nt, nX, dt, dX, Tl, Tr, rvalue)
	//exact(nX, float64(alpha), float64(maxtime), float64(dX), float64(1.0), float64(Tr), float64(Tl))
	//DuFort_Frankel_Explicit_Scheme(nt, nX, dX, dt, maxtime, rvalue, Tr, Tl)

}
