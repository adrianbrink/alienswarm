package swarm

import (
  "testing"
  "os"
)

func benchmark(numAliens int, numIterations int, b *testing.B) {
  os.Stdout, _ = os.Open(os.DevNull)
  sim := NewSim("../sample/input.50x50", numAliens, numIterations)
  b.ResetTimer()
  for n := 0; n < b.N; n++ {
    sim.Run()
  }
}

func BenchmarkSwarm1000(b *testing.B) {
  benchmark(1000, 100000, b)
}

func BenchmarkSwarm100(b *testing.B) {
  benchmark(100, 100000, b)
}

func BenchmarkSwarm10(b *testing.B) {
  benchmark(10, 100000, b)
}

func BenchmarkSwarm2(b *testing.B) {
  benchmark(2, 100000, b)
}

func BenchmarkSwarm1(b *testing.B) {
  benchmark(1, 100000, b)
}
