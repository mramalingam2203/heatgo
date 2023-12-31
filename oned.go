package main

/*
func Laasonen_Simple_Implicit_Scheme(numerical *[][]float64, delta_x float64, delta_t float64, output_time float64, nodes int, r float64, T_sur float64, T_init float64) {

	for n := 0; n <= int(output_time/delta_t); n++ {
		numerical[0][n] = T_sur
		numerical[nodes-1][n] = T_sur
	}

	for i := 1; i < (nodes - 1); i++ {
		numerical[i][0] = T_init
	}

	lower_diag := make([]float64, nodes-2)
	upper_diag := make([]float64, nodes-2)
	main_diag := make([]float64, nodes-2)
	b := make([]float64, nodes-2)

	for n := 0; n < int(output_time/delta_t); n++ {
		for k := 0; k < nodes-3; k++ {
			upper_diag[k] = -r
		}

		for k := 0; k < nodes-2; k++ {
			if k == 0 {
				b[k] = numerical[k+1][n] + r*T_sur
			} else if k == nodes-3 {
				b[k] = numerical[k+1][n] + r*T_sur
			} else {
				b[k] = numerical[k+1][n]
			}
		}

		TDMA_Solver(lower_diag, main_diag, upper_diag, b, nodes)
		for i := 0; i < nodes-2; i++ {
			numerical[i+1][n+1] = b[i]
		}
	}

}
*/
func Crank_Nicholson_Implicit_Scheme(numerical *[][]float64, delta_x float64, delta_t float64, output_time float64, nodes int, r float64, T_sur float64, T_init float64) {

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
