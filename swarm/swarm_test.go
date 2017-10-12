package swarm

import (
	"os"
	"testing"
)

func benchmark(size string, numAliens int, b *testing.B) {
	os.Stdout, _ = os.Open(os.DevNull)
	sim := NewSim("../sample/input."+size, numAliens, 100000)
	for n := 0; n < b.N; n++ {
		sim.Step()
	}
}

func BenchmarkSwarm100x100x1000(b *testing.B) {
	benchmark("100x100", 1000, b)
}
func BenchmarkSwarm100x100x500(b *testing.B) {
	benchmark("100x100", 500, b)
}
func BenchmarkSwarm100x100x100(b *testing.B) {
	benchmark("100x100", 100, b)
}

func BenchmarkSwarm50x50x1000(b *testing.B) {
	benchmark("50x50", 1000, b)
}

func BenchmarkSwarm50x50x100(b *testing.B) {
	benchmark("50x50", 100, b)
}

func BenchmarkSwarm50x50x10(b *testing.B) {
	benchmark("50x50", 10, b)
}

func BenchmarkSwarm50x50x2(b *testing.B) {
	benchmark("50x50", 2, b)
}

func BenchmarkSwarm50x50x1(b *testing.B) {
	benchmark("50x50", 1, b)
}
