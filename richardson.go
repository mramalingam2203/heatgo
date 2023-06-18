package main

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
