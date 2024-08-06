package main

import (
	"testing"
	"unsafe"
)

func BenchmarkArrayConversion(b *testing.B) {
	sizes := []struct {
		rows, cols int
	}{
		{320, 200},
		{640, 400},
		{960, 600},
	}

	for _, size := range sizes {
		b.Run("UnsafeConversion", func(b *testing.B) {
			mat := make([][]float64, size.rows)
			for i := range mat {
				mat[i] = make([]float64, size.cols)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				_ = (*[1 << 30]float64)(unsafe.Pointer(&mat[0][0]))[:size.rows*size.cols]
			}
		})

		b.Run("CopyConversion", func(b *testing.B) {
			chunks := make([][]byte, size.rows)
			for i := range chunks {
				chunks[i] = make([]byte, size.cols)
			}
			b.ResetTimer()
			for i := 0; i < b.N; i++ {
				l := 0
				for _, v := range chunks {
					l += len(v)
				}
				result := make([]byte, 0, l)
				for _, v := range chunks {
					result = append(result, v...)
				}
			}
		})
	}
}
