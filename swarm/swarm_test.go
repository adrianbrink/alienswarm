package swarm

import "testing"

func BenchmarkSwarm10(b *testing.B) {
  sim := NewSim("../input.txt", 10, 1000)
  b.ResetTimer()
	for n := 0; n < b.N; n++ {
    sim.Run()
	}
}

func BenchmarkSwarm2(b *testing.B) {
  sim := NewSim("../input.txt", 2, 1000)
  b.ResetTimer()
	for n := 0; n < b.N; n++ {
    sim.Run()
	}
}

func BenchmarkSwarm1(b *testing.B) {
  sim := NewSim("../input.txt", 1, 1000)
  b.ResetTimer()
	for n := 0; n < b.N; n++ {
    sim.Run()
	}
}
