package main

import (
	"aoc2024/day1"
	"testing"
)

func BenchmarkDay1_2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		day1.Second([]string{
			"3   4",
			"4   3",
			"2   5",
			"1   3",
			"3   9",
			"3   3",
		})
        b.Logf("Elapsed: %d microseconds", b.Elapsed().Microseconds())
	}
    b.ReportAllocs()
}
