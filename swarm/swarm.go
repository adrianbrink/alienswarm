package swarm

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

type Sim struct {
	World            *World
	Aliens           []*Alien
	MaxIterations    int
	CurrentIteration int
}

func (sim *Sim) Step() {
	// Move all aliens.
	var wg sync.WaitGroup
	wg.Add(len(sim.Aliens))
	for _, a := range sim.Aliens {
		go func(alien *Alien) {
			defer wg.Done()
			alien.RandomWalk()
		}(a)
	}
	// Block until all Aliens have updated their state.
	wg.Wait()

	// Manipulating the slice is too expensive, we just copy over
	// the aliens that survived and reset the sim state.
	alienSurvived := make([]*Alien, 0)

	// We want to parallelize the
	survivorChan := make(chan *Alien, 1)
	destroyedChan := make(chan string, 1)

	// Aggregate survives/destroys into a chan
	for _, a := range sim.Aliens {
		go func(alien *Alien) {
			if alien.City.NumVisitors <= 1 {
				survivorChan <- alien
			} else {
				// This allows us to know the size of chan ahead of time.
				destroyedChan <- alien.City.Name
			}
		}(a)
	}

	// Update Sim State
	for i := 0; i < len(sim.Aliens); i++ {
		select {
		case alien := <-survivorChan:
			alienSurvived = append(alienSurvived, alien)
		case cityName := <-destroyedChan:
			city, found := sim.World.FindCity(cityName)
			// The destroyedChan will contain duplicate cities, this is fine, we just skip after the first delete.
			if found {
				aliens := city.GetVisitors()
				sim.World.DeleteCity(cityName)

				// Builds up a string for each alien, handling the edge case at the end.
				numAliens := len(aliens)
				names := make([]string, numAliens)
				for i, a := range aliens {
					names[i] = a.Name
				}
				killList := strings.Join(names[:numAliens-1], ", ")
				killList += " and " + aliens[numAliens-1].Name
				fmt.Printf("%s was destroyed by aliens %s!\n", city, killList)
			}
		}
	}
	sim.Aliens = alienSurvived
}

func (sim *Sim) Run() {
	for sim.CurrentIteration <= sim.MaxIterations && len(sim.Aliens) != 0 {
		fmt.Println("Stepping", sim.CurrentIteration, len(sim.Aliens))
		sim.Step()
		sim.CurrentIteration++
	}
	fmt.Println("Done!")
}

func NewSim(cityFile string, numAliens int, maxIterations int) *Sim {
	s := &Sim{MaxIterations: maxIterations}
	s.World = BuildWorld(cityFile)
	s.Aliens = make([]*Alien, numAliens)
	for i := 0; i < numAliens; i++ {
		s.Aliens[i] = NewAlien(strconv.Itoa(i), s.World.RandomCity())
	}
	s.CurrentIteration = 0
	return s
}
