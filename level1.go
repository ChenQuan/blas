package goblas

import (
	"sync"
)

// Snrm2 computes the Euclidean norm of a vector,
//  sqrt(\sum_i x[i] * x[i]).
// This function returns 0 if incX is negative.
func Snrm2(N int, x []float32, incX int) float32 {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return 0
	}
	if incX > 0 && (N-1)*incX >= len(x) {
		panic(badX)
	}
	if N < 2 {
		if N == 1 {
			return Abs(x[0])
		}
		if N == 0 {
			return 0
		}
		if N < 1 {
			panic(negativeN)
		}
	}
	var (
		scale      float32 = 0
		sumSquares float32 = 1
	)

	if incX == 1 {
		x = x[:N]
		var wg sync.WaitGroup
		for _, v := range x {
			if v == 0 {
				continue
			}
			if IsNaN(v) {
				return NaN()
			}
			wg.Add(1)
			go func(v float32) {
				defer wg.Done()
				absx := Abs(v)

				if scale < absx {
					sumSquares = 1 + sumSquares*(scale/absx)*(scale/absx)
					scale = absx
				} else {
					sumSquares = sumSquares + (absx/scale)*(absx/scale)
				}
			}(v)

		}
		wg.Wait()

		if IsInf(scale, 1) {
			return Inf(1)
		}
		return scale * Sqrt(sumSquares)
	}
	var wg sync.WaitGroup

	for ix := 0; ix < N*incX; ix += incX {
		v := x[ix]
		if v == 0 {
			continue
		}
		absxi := Abs(v)
		if IsNaN(absxi) {
			return NaN()
		}
		wg.Add(1)

		go func(v float32) {
			defer wg.Done()
			absx := Abs(v)

			if scale < absx {
				sumSquares = 1 + sumSquares*(scale/absx)*(scale/absx)
				scale = absx
			} else {
				sumSquares = sumSquares + (absx/scale)*(absx/scale)
			}
		}(v)
	}
	wg.Wait()
	if IsInf(scale, 1) {
		return Inf(1)
	}
	return scale * Sqrt(sumSquares)
}
