func Hopscotch(nt int, nx int, ny int, delta_t float64, delta_x float64, delta_y float64,
	output_time float64, r float64, T_sur float64, T_init float64) {

	for i := 0; i < nx; i++ {
		for j := 0; j < ny; j++ {
			T[0][x][y] = T_init
		}
	}

	for t := 0; t < ntime; t++ {
		for i := 0; i < nx; i++ {
			T[t][0][i] = T_west
		}
	}

	for t := 0; t < ntime; t++ {
		for j := 0; j < ny; j++ {
			T[t][nx][y] = T_east
		}
	}

	for t := 0; t < ntime; t++ {
		for i := 0; j < ny; j++ {
			T[t][nx][y] = T_south
		}
	}

	for t := 0; t < ntime; t++ {
		for j := 0; j < ny; j++ {
			T[t][nx][y] = T_north
		}
	}

}
