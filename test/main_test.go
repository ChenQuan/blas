package test

import (
	"fmt"
	"testing"
)

const N = 1000000

// 常规 for 迭代 slice
func ForSlice(s []string) {

	sLen := len(s)
	for i := 0; i < sLen; i++ {

		_, b := i, s[i]
		var a float64
		a = 45.4545566565656
		a = a * a
		a = a * a
		a = a * a
		a = a * a

		if b == "" {
			fmt.Println(a)
		}

	}

}

// for range 迭代 slice
func RangeForSlice(s []string) {

	for i, v := range s {

		_, b := i, v

		var a float64
		a = 45.4545566565656
		a = a * a
		a = a * a
		a = a * a
		a = a * a

		if b == "" {
			fmt.Println(a)
		}

	}

}

// for range 迭代 slice
func RangeForSliceWithI(s []string) {

	for i, _ := range s {

		_, b := i, s[i]

		var a float64
		a = 45.4545566565656
		a = a * a
		a = a * a
		a = a * a
		a = a * a

		if b == "" {
			fmt.Println(a)
		}
	}

}

// 初始化 slice
func initSlice() []string {
	s := make([]string, N)

	for i := 0; i < N; i++ {
		s[i] = "www.flysnow.org"
	}
	return s
}

// 基准测试函数
func BenchmarkForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSlice(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}
func BenchmarkRangeForSliceWithI(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSliceWithI(s)
	}
}

//pkg: github.com/yunqi/blas/test
//BenchmarkForSlice-8             	    5000	    292603 ns/op
//BenchmarkRangeForSlice-8        	    1000	   1418406 ns/op
//BenchmarkRangeForSliceWithI-8   	    5000	    297599 ns/op

//enchmarkForSlice-8             	       1	77737869442 ns/op
//BenchmarkRangeForSlice-8        	       1	81947644536 ns/op
//BenchmarkRangeForSliceWithI-8   	       1	77840163906 ns/op

//BenchmarkForSlice-8             	       1	311246656114 ns/op
//BenchmarkRangeForSlice-8        	       1	302468630142 ns/op
//BenchmarkRangeForSliceWithI-8   	       1	301502139560 ns/op
