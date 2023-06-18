package main

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
