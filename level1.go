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

// Sasum computes the sum of the absolute values of the elements of x.
//  \sum_i |x[i]|
// Sasum returns 0 if incX is negative.
func Sasum(N int, x []float32, incX int) float32 {

	if N < 0 {
		panic(negativeN)
	}
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return 0
	}
	if incX > 0 && (N-1)*incX >= len(x) {
		panic(badX)
	}
	var sum float32
	if incX == 1 {
		x = x[:N]
		var wg sync.WaitGroup
		for _, v := range x {
			wg.Add(1)
			go func(v float32) {
				defer wg.Done()
				sum += Abs(v)
			}(v)
		}
		wg.Wait()
		return sum
	}
	var wg sync.WaitGroup

	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(v int) {
			defer wg.Done()
			sum += Abs(x[i*incX])
		}(i)
	}
	wg.Wait()
	return sum
}

// Isamax returns the index of an element of x with the largest absolute value.
// If there are multiple such indices the earliest is returned.
// Isamax returns -1 if n == 0.
// 未完成
func Isamax(N int, x []float32, incX int) int {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return -1
	}
	if incX > 0 && (N-1)*incX >= len(x) {
		panic(badX)
	}
	if N < 2 {
		if N == 1 {
			return 0
		}
		if N == 0 {
			return -1 // Netlib returns invalid index when n == 0
		}
		if N < 1 {
			panic(negativeN)
		}
	}

	var idx int
	max := Abs(x[0])
	if incX == 1 {
		for i, v := range x[:N] {
			absV := Abs(v)
			if absV > max {
				max = absV
				idx = i
			}
		}
		return idx
	}
	ix := incX
	for i := 1; i < N; i++ {
		v := x[ix]
		absV := Abs(v)
		if absV > max {
			max = absV
			idx = i
		}
		ix += incX
	}
	return idx
}
func Sswap(N int, x []float32, incX int, y []float32, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if N < 1 {
		if N == 0 {
			return
		}
		panic(negativeN)
	}
	if (incX > 0 && (N-1)*incX >= len(x)) || (incX < 0 && (1-N)*incX >= len(x)) {
		panic(badX)
	}
	if (incY > 0 && (N-1)*incY >= len(y)) || (incY < 0 && (1-N)*incY >= len(y)) {
		panic(badY)
	}
	if incX == 1 && incY == 1 {
		x = x[:N]
		var wg sync.WaitGroup
		for i, v := range x {
			wg.Add(1)
			go func(i int, v float32) {
				defer wg.Done()

				x[i], y[i] = y[i], v
			}(i, v)
		}
		wg.Wait()
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-N + 1) * incX
	}
	if incY < 0 {
		iy = (-N + 1) * incY
	}
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(iy, ix int) {
			defer wg.Done()
			x[ix], y[iy] = y[iy], x[ix]
		}(iy, ix)

		ix += incX
		iy += incY

	}
	wg.Wait()
}

// Scopy copies the elements of x into the elements of y.
//  y[i] = x[i] for all i
//
// Float32 implementations are autogenerated and not directly tested.
func Scopy(N int, x []float32, incX int, y []float32, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if N < 1 {
		if N == 0 {
			return
		}
		panic(negativeN)
	}
	if (incX > 0 && (N-1)*incX >= len(x)) || (incX < 0 && (1-N)*incX >= len(x)) {
		panic(badX)
	}
	if (incY > 0 && (N-1)*incY >= len(y)) || (incY < 0 && (1-N)*incY >= len(y)) {
		panic(badY)
	}
	if incX == 1 && incY == 1 {
		copy(y[:N], x[:N])
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-N + 1) * incX
	}
	if incY < 0 {
		iy = (-N + 1) * incY
	}
	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)

		go func(iy, ix int) {
			defer wg.Done()
			y[iy] = x[ix]
		}(iy, ix)
		ix += incX
		iy += incY
	}
	wg.Wait()
}

// Saxpy adds alpha times x to y
//  y[i] += alpha * x[i] for all i

func Saxpy(N int, alpha float32, x []float32, incX int, y []float32, incY int) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if N < 1 {
		if N == 0 {
			return
		}
		panic(negativeN)
	}
	if (incX > 0 && (N-1)*incX >= len(x)) || (incX < 0 && (1-N)*incX >= len(x)) {
		panic(badX)
	}
	if (incY > 0 && (N-1)*incY >= len(y)) || (incY < 0 && (1-N)*incY >= len(y)) {
		panic(badY)
	}
	if alpha == 0 {
		return
	}

	if incX == 1 && incY == 1 {
		if len(x) < N {
			panic(badLenX)
		}
		if len(y) < N {
			panic(badLenY)
		}
		x = x[:N]
		var wg sync.WaitGroup
		for i, v := range x {
			wg.Add(1)
			go func(i int, v float32) {
				defer wg.Done()
				y[i] += alpha * v
			}(i, v)
		}
		wg.Wait()
		return
	}
	var ix, iy int
	if incX < 0 {
		ix = (-N + 1) * incX
	}
	if incY < 0 {
		iy = (-N + 1) * incY
	}
	if ix >= len(x) || ix+(N-1)*incX >= len(x) {
		panic(badLenX)
	}
	if iy >= len(y) || iy+(N-1)*incY >= len(y) {
		panic(badLenY)
	}

	var wg sync.WaitGroup
	for i := 0; i < N; i++ {
		wg.Add(1)

		go func(ix int, iy int) {
			defer wg.Done()
			y[iy] += alpha * x[ix]
		}(ix, iy)
		ix += incX
		iy += incY
	}
	wg.Wait()
}

// Srotg computes the plane rotation
//   _    _      _ _       _ _
//  |  c s |    | a |     | r |
//  | -s c |  * | b |   = | 0 |
//   ‾    ‾      ‾ ‾       ‾ ‾
// where
//  r = ±√(a^2 + b^2)
//  c = a/r, the cosine of the plane rotation
//  s = b/r, the sine of the plane rotation
//
// NOTE: There is a discrepancy between the refence implementation and the BLAS
// technical manual regarding the sign for r when a or b are zero.
// Srotg agrees with the definition in the manual and other
// common BLAS implementations.

func Srotg(a, b float32) (c, s, r, z float32) {
	if b == 0 && a == 0 {
		return 1, 0, a, 0
	}
	absA := Abs(a)
	absB := Abs(b)
	aGTb := absA > absB
	r = Hypot(a, b)
	if aGTb {
		r = Copysign(r, a)
	} else {
		r = Copysign(r, b)
	}
	c = a / r
	s = b / r
	if aGTb {
		z = s
	} else if c != 0 { // r == 0 case handled above
		z = 1 / c
	} else {
		z = 1
	}
	return
}
