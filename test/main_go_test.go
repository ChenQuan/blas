package test

import (
	"fmt"
	"sync"
	"testing"
)

// 常规 for 迭代 slice
func ForSliceGo(s []string) {
	var wg sync.WaitGroup
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
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
		}(i)

	}
	wg.Wait()

}

// for range 迭代 slice
func RangeForSliceGo(s []string) {
	var wg sync.WaitGroup

	for i, v := range s {
		wg.Add(1)
		go func(i int, v string) {
			defer wg.Done()
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
		}(i, v)
	}
	wg.Wait()

}

// for range 迭代 slice
func RangeForSliceWithIGo(s []string) {
	var wg sync.WaitGroup

	for i, _ := range s {

		wg.Add(1)

		go func(i int) {
			defer wg.Done()

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
		}(i)

	}
	wg.Wait()

}

// 基准测试函数
func BenchmarkForSliceGo(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ForSlice(s)
	}
}

func BenchmarkRangeForSliceGo(b *testing.B) {
	s := initSlice()

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		RangeForSlice(s)
	}
}
func BenchmarkRangeForSliceWithIGo(b *testing.B) {
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
