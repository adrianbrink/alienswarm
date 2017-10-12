package swarm

import (
	"fmt"
	"strconv"
	"strings"
)

type Sim struct {
	World            *World
	Aliens           []*Alien
	MaxIterations    int
	CurrentIteration int
}

func (sim *Sim) Step() {
	// Move all aliens and build up a map of city -> aliens
	occupancy := make(map[string][]*Alien)
	for _, a := range sim.Aliens {
		a.RandomWalk()
		name := a.City.Name
		aliens, found := occupancy[name]
		if !found {
			aliens = []*Alien{a}
			occupancy[name] = aliens
		} else {
			occupancy[name] = append(aliens, a)
		}
	}

	// Manipulating the slice is too expensive, we just copy over
	// the aliens that survived and reset the sim state.
	alienSurvived := make([]*Alien, 0)

	// For the occupancy map, if len(aliens) > 1, destroy everything.
	for city, aliens := range occupancy {
		numAliens := len(aliens)
		if numAliens > 1 {
			sim.World.DeleteCity(city)

			// Builds up a string for each alien, handling the edge case at the end.
			names := make([]string, numAliens)
			for i, a := range aliens {
				names[i] = a.Name
			}
			killList := strings.Join(names[:numAliens-1], ", ")
			killList += " and " + aliens[numAliens-1].Name
			fmt.Printf("%s was destroyed by aliens %s!\n", city, killList)
		} else {
			alienSurvived = append(alienSurvived, aliens[0])
		}
	}

	sim.Aliens = alienSurvived
}

func (sim *Sim) Run() {
	for sim.CurrentIteration <= sim.MaxIterations && len(sim.Aliens) != 0 {
		sim.Step()
		sim.CurrentIteration++
	}
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
