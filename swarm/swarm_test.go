package swarm

import "testing"

func BenchmarkSwarm1000(b *testing.B) {
	sim := NewSim("../sample/input.50x50", 1000, 100000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}

func BenchmarkSwarm100(b *testing.B) {
	sim := NewSim("../sample/input.50x50", 100, 100000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}

func BenchmarkSwarm10(b *testing.B) {
	sim := NewSim("../sample/input.50x50", 10, 100000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}

func BenchmarkSwarm2(b *testing.B) {
	sim := NewSim("../sample/input.50x50", 2, 100000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}

func BenchmarkSwarm1(b *testing.B) {
	sim := NewSim("../sample/input.50x50", 1, 100000)
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		sim.Run()
	}
}
