package structs

import "testing"

// Results:
// Benchmark_SingleArrayAdd-12    	   56902	    153257 ns/op	  459201 B/op	       1 allocs/op
// Benchmark_Vector2ArrayAdd-12    	  110671	    225833 ns/op	  444712 B/op	       0 allocs/op
// Benchmark_Factor5ArrayAdd-12    	69872607	    55.55 ns/op	      87 B/op	           0 allocs/op

//go:noinline
func Benchmark_SingleArrayAdd(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	array := NewSingleArray[string]()

	for N := 0; N < b.N; N++ {
		array.Add("a")
	}
}

//go:noinline
func Benchmark_Vector2ArrayAdd(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	array := NewVectorArray[string](2)

	for N := 0; N < b.N; N++ {
		array.Add("a")
	}
}

//go:noinline
func Benchmark_Factor5ArrayAdd(b *testing.B) {
	b.ResetTimer()
	b.ReportAllocs()

	array := NewFactorArray[string](5)

	for N := 0; N < b.N; N++ {
		array.Add("a")
	}
}
