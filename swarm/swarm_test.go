package swarm

import (
	"os"
	"testing"
)

func benchmark(size string, numAliens int, numIterations int, b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	sim := NewSim("../sample/input."+size, numAliens, numIterations)
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}

func BenchmarkSwarm100x100x1000(b *testing.B) {
	benchmark("100x100", 1000, 10000, b)
}
func BenchmarkSwarm100x100x500(b *testing.B) {
	benchmark("100x100", 500, 10000, b)
}
func BenchmarkSwarm100x100x100(b *testing.B) {
	benchmark("100x100", 100, 10000, b)
}

func BenchmarkSwarm50x50x1000(b *testing.B) {
	benchmark("50x50", 1000, 10000, b)
}

func BenchmarkSwarm50x50x100(b *testing.B) {
	benchmark("50x50", 100, 10000, b)
}

func BenchmarkSwarm50x50x10(b *testing.B) {
	benchmark("50x50", 10, 10000, b)
}

func BenchmarkSwarm50x50x2(b *testing.B) {
	benchmark("50x50", 2, 100000, b)
}

func BenchmarkSwarm50x50x1(b *testing.B) {
	benchmark("50x50", 1, 100000, b)
}
