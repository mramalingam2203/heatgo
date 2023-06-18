package main

import "fmt"

func DuFort_Frankel_Explicit_Scheme(nt int, nodes int, delta_x float64, delta_t float64, output_time float64, r float64, T_sur float64, T_init float64) {

	numerical := make([][]float64, nt)
	for i := 0; i < nt; i++ {
		numerical[i] = make([]float64, nodes)
	}

	for n := 0; n < nodes; n++ {
		numerical[0][n] = T_init      // BC at 0 and all n (node #0)
		numerical[nodes-1][n] = T_sur // BC at 31 and all n (node #last_node)
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #1
		numerical[i][0] = T_init // IC at n = 0, for all i, except node node i = 0 and i = nodes
	}

	for i := 1; i < int(nodes-1); i++ { // Initial Conditions #2
		numerical[i][1] = r*numerical[i-1][0] + (1.0-(2.0*r))*numerical[i][0] + r*numerical[i+1][0] // IC at n = 1, for all i, except node node i = 0 and i = nodes
	}

	for i := 1; i < nt-1; i++ {
		for n := 1; n < nodes-1; n++ {
			numerical[i][n+1] = ((2.0*r)/(1.0+2.0*r))*numerical[i-1][n] + ((2.0*r)/(1.0+2.0*r))*numerical[i+1][n] + ((1.0-2.0*r)/(1.0+2.0*r))*numerical[i][n-1]
		}
	}

	fmt.Println(numerical)

}
