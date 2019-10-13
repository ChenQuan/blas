package goblas

import (
	"math"
	"sync"
)

// Double

// Dnrm2 computes the Euclidean norm of a vector,
//  sqrt(\sum_i x[i] * x[i]).
// This function returns 0 if incX is negative.
// 未完成，有问题
func Dnrm2(N int, x []float64, incX int) float64 {
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
			return math.Abs(x[0])
		}
		if N == 0 {
			return 0
		}
		if N < 1 {
			panic(negativeN)
		}
	}

	var sum float64
	var wg sync.WaitGroup
	if incX == 1 {
		x = x[:N]
		var wg sync.WaitGroup
		for _, v := range x {
			if v == 0 {
				continue
			}
			if math.IsNaN(v) {
				return math.NaN()
			}
			wg.Add(1)
			go func(v float64) {
				defer wg.Done()
				absx := math.Abs(v)
				sum += absx * absx
			}(v)

		}
		wg.Wait()

		if math.IsInf(sum, 1) {
			return math.Inf(1)
		}
		return math.Sqrt(sum)
	}
	for ix := 0; ix < N*incX; ix += incX {
		v := x[ix]
		if v == 0 {
			continue
		}
		absxi := math.Abs(v)
		if math.IsNaN(absxi) {
			return math.NaN()
		}
		wg.Add(1)

		go func(v float64) {
			defer wg.Done()
			absx := math.Abs(v)

			sum += absx * absx
		}(v)
	}
	wg.Wait()

	if math.IsInf(sum, 1) {
		return math.Inf(1)
	}
	return math.Sqrt(sum)
}

// Dasum computes the sum of the absolute values of the elements of x.
//  \sum_i |x[i]|
// Dasum returns 0 if incX is negative.
func Dasum(N int, x []float64, incX int) float64 {
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
	var wg sync.WaitGroup
	var sum float64

	if incX == 1 {
		x = x[:N]

		for _, v := range x {
			wg.Add(1)

			go func(v float64) {
				defer wg.Done()
				sum += math.Abs(v)
			}(v)
		}
		wg.Wait()

		return sum
	}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			sum += math.Abs(x[i*incX])
		}(i)
	}
	wg.Wait()
	return sum
}

// Idamax returns the index of an element of x with the largest absolute value.
// If there are multiple such indices the earliest is returned.
// Idamax returns -1 if n == 0.
func Idamax(N int, x []float64, incX int) int {
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
	idx := 0
	max := math.Abs(x[0])
	if incX == 1 {
		for i, v := range x[:N] {
			absV := math.Abs(v)
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
		absV := math.Abs(v)
		if absV > max {
			max = absV
			idx = i
		}
		ix += incX
	}
	return idx
}

// Dswap exchanges the elements of two vectors.
//  x[i], y[i] = y[i], x[i] for all i
func Dswap(N int, x []float64, incX int, y []float64, incY int) {
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
	var wg sync.WaitGroup

	if incX == 1 && incY == 1 {
		x = x[:N]
		for i, v := range x {
			wg.Add(1)
			go func(i int, v float64) {
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

// Dcopy copies the elements of x into the elements of y.
//  y[i] = x[i] for all i
func Dcopy(N int, x []float64, incX int, y []float64, incY int) {
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

// Daxpy adds alpha times x to y
//  y[i] += alpha * x[i] for all i
func Daxpy(N int, alpha float64, x []float64, incX int, y []float64, incY int) {
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
	var wg sync.WaitGroup

	if incX == 1 && incY == 1 {
		if len(x) < N {
			panic(badLenX)
		}
		if len(y) < N {
			panic(badLenY)
		}
		x = x[:N]
		for i, v := range x {
			wg.Add(1)
			go func(i int, v float64) {
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

// Drotg computes the plane rotation
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
// Drotg agrees with the definition in the manual and other
// common BLAS implementations.
func Drotg(a, b float64) (c, s, r, z float64) {
	if b == 0 && a == 0 {
		return 1, 0, a, 0
	}
	absA := math.Abs(a)
	absB := math.Abs(b)
	aGTb := absA > absB
	r = math.Hypot(a, b)
	if aGTb {
		r = math.Copysign(r, a)
	} else {
		r = math.Copysign(r, b)
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

// Drotmg computes the modified Givens rotation. See
// http://www.netlib.org/lapack/explore-html/df/deb/drotmg_8f.html
// for more details.
func Drotmg(d1, d2, x1, y1 float64) (p DrotmParams, rd1, rd2, rx1 float64) {
	var p1, p2, q1, q2, u float64

	const (
		gam    = 4096.0
		gamsq  = 16777216.0
		rgamsq = 5.9604645e-8
	)

	if d1 < 0 {
		p.Flag = Rescaling
		return
	}

	p2 = d2 * y1
	if p2 == 0 {
		p.Flag = Identity
		rd1 = d1
		rd2 = d2
		rx1 = x1
		return
	}
	p1 = d1 * x1
	q2 = p2 * y1
	q1 = p1 * x1

	absQ1 := math.Abs(q1)
	absQ2 := math.Abs(q2)

	if absQ1 < absQ2 && q2 < 0 {
		p.Flag = Rescaling
		return
	}

	if d1 == 0 {
		p.Flag = Diagonal
		p.H[0] = p1 / p2
		p.H[3] = x1 / y1
		u = 1 + p.H[0]*p.H[3]
		rd1, rd2 = d2/u, d1/u
		rx1 = y1 / u
		return
	}

	// Now we know that d1 != 0, and d2 != 0. If d2 == 0, it would be caught
	// when p2 == 0, and if d1 == 0, then it is caught above

	if absQ1 > absQ2 {
		p.H[1] = -y1 / x1
		p.H[2] = p2 / p1
		u = 1 - p.H[2]*p.H[1]
		rd1 = d1
		rd2 = d2
		rx1 = x1
		p.Flag = OffDiagonal
		// u must be greater than zero because |q1| > |q2|, so check from netlib
		// is unnecessary
		// This is left in for ease of comparison with complex routines
		//if u > 0 {
		rd1 /= u
		rd2 /= u
		rx1 *= u
		//}
	} else {
		p.Flag = Diagonal
		p.H[0] = p1 / p2
		p.H[3] = x1 / y1
		u = 1 + p.H[0]*p.H[3]
		rd1 = d2 / u
		rd2 = d1 / u
		rx1 = y1 * u
	}

	for rd1 <= rgamsq || rd1 >= gamsq {
		if p.Flag == OffDiagonal {
			p.H[0] = 1
			p.H[3] = 1
			p.Flag = Rescaling
		} else if p.Flag == Diagonal {
			p.H[1] = -1
			p.H[2] = 1
			p.Flag = Rescaling
		}
		if rd1 <= rgamsq {
			rd1 *= gam * gam
			rx1 /= gam
			p.H[0] /= gam
			p.H[2] /= gam
		} else {
			rd1 /= gam * gam
			rx1 *= gam
			p.H[0] *= gam
			p.H[2] *= gam
		}
	}

	for math.Abs(rd2) <= rgamsq || math.Abs(rd2) >= gamsq {
		if p.Flag == OffDiagonal {
			p.H[0] = 1
			p.H[3] = 1
			p.Flag = Rescaling
		} else if p.Flag == Diagonal {
			p.H[1] = -1
			p.H[2] = 1
			p.Flag = Rescaling
		}
		if math.Abs(rd2) <= rgamsq {
			rd2 *= gam * gam
			p.H[1] /= gam
			p.H[3] /= gam
		} else {
			rd2 /= gam * gam
			p.H[1] *= gam
			p.H[3] *= gam
		}
	}
	return
}

// Drot applies a plane transformation.
//  x[i] = c * x[i] + s * y[i]
//  y[i] = c * y[i] - s * x[i]
func Drot(N int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
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
	var wg sync.WaitGroup

	if incX == 1 && incY == 1 {
		x = x[:N]
		for i, vx := range x {
			wg.Add(1)
			go func(i int, vx float64) {
				defer wg.Done()
				vy := y[i]
				x[i], y[i] = c*vx+s*vy, c*vy-s*vx
			}(i, vx)
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
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(ix, iy int) {
			defer wg.Done()
			vx := x[ix]
			vy := y[iy]
			x[ix], y[iy] = c*vx+s*vy, c*vy-s*vx
		}(ix, iy)
		ix += incX
		iy += incY
	}
	wg.Wait()
}

// Drotm applies the modified Givens rotation to the 2×n matrix.
func Drotm(N int, x []float64, incX int, y []float64, incY int, p DrotmParams) {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if N <= 0 {
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

	var h11, h12, h21, h22 float64
	var ix, iy int
	switch p.Flag {
	case Identity:
		return
	case Rescaling:
		h11 = p.H[0]
		h12 = p.H[2]
		h21 = p.H[1]
		h22 = p.H[3]
	case OffDiagonal:
		h11 = 1
		h12 = p.H[2]
		h21 = p.H[1]
		h22 = 1
	case Diagonal:
		h11 = p.H[0]
		h12 = 1
		h21 = -1
		h22 = p.H[3]
	}
	if incX < 0 {
		ix = (-N + 1) * incX
	}
	if incY < 0 {
		iy = (-N + 1) * incY
	}
	var wg sync.WaitGroup

	if incX == 1 && incY == 1 {
		x = x[:N]
		for i, vx := range x {
			wg.Add(1)
			go func(i int, vx float64) {
				defer wg.Done()
				vy := y[i]
				x[i], y[i] = vx*h11+vy*h12, vx*h21+vy*h22
			}(i, vx)
		}
		wg.Wait()
		return
	}
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func(ix, iy int) {
			defer wg.Done()
			vx := x[ix]
			vy := y[iy]
			x[ix], y[iy] = vx*h11+vy*h12, vx*h21+vy*h22
		}(ix, iy)
		ix += incX
		iy += incY
	}
	wg.Wait()
	return
}

// Dscal scales x by alpha.
//  x[i] *= alpha
// Dscal has no effect if incX < 0.
func Dscal(N int, alpha float64, x []float64, incX int) {
	if incX < 1 {
		if incX == 0 {
			panic(zeroIncX)
		}
		return
	}
	if (N-1)*incX >= len(x) {
		panic(badX)
	}
	if N < 1 {
		if N == 0 {
			return
		}
		panic(negativeN)
	}
	var wg sync.WaitGroup

	if alpha == 0 {
		if incX == 1 {
			x = x[:N]
			for i := range x {
				wg.Add(1)
				go func(i int) {
					defer wg.Done()
					x[i] = 0
				}(i)
			}
			wg.Wait()
			return
		}
		for ix := 0; ix < N*incX; ix += incX {
			wg.Add(1)
			go func(ix int) {
				defer wg.Done()
				x[ix] = 0
			}(ix)
		}
		wg.Wait()
		return
	}
	if incX == 1 {
		x = x[:N]
		for i := range x {
			wg.Add(1)
			go func(i int) {
				defer wg.Done()
				x[i] *= alpha
			}(i)
		}
		wg.Wait()
		return
	}

	for ix := 0; ix < N*incX; ix += incX {
		wg.Add(1)
		go func(ix int) {
			defer wg.Done()
			x[ix] *= alpha
		}(ix)
	}
	wg.Wait()
}

// Ddot computes the dot product of the two vectors
//  \sum_i x[i]*y[i]
func Ddot(N int, x []float64, incX int, y []float64, incY int) float64 {
	if incX == 0 {
		panic(zeroIncX)
	}
	if incY == 0 {
		panic(zeroIncY)
	}
	if N <= 0 {
		if N == 0 {
			return 0
		}
		panic(negativeN)
	}
	var wg sync.WaitGroup
	var sum float64
	if incX == 1 && incY == 1 {
		if len(x) < N {
			panic(badLenX)
		}
		if len(y) < N {
			panic(badLenY)
		}
		x = x[:N]
		for i, v := range x {
			wg.Add(1)
			go func(i int, v float64) {
				wg.Done()
				sum += y[i] * v
			}(i, v)
		}
		wg.Wait()
		return sum
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
	for i := 0; i < N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			sum += x[ix] * y[iy]
		}()
		ix += incX
		iy += incY
	}
	wg.Wait()
	return sum
}
