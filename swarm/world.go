package swarm

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strings"
)

// World is set of helper functions around the City graph.
type World struct {
	cities map[string]*City
}

func (w *World) PrintWorld() {
	for _, v := range w.cities {
		fmt.Println(v)
	}
}

func (w *World) FindCity(name string) (*City, bool) {
	city, found := w.cities[name]
	return city, found
}

func (w *World) RandomCity() *City {
	choice := rand.Intn(len(w.cities))
	count := 0
	// The order here is undefined for `range` -- we randomly choose an index
	// but if the order happens to have some relationship with rand.Intn then
	// this could be come deterministic.
	for _, v := range w.cities {
		if count == choice {
			return v
		}
		count++
	}
	return nil
}

func (w *World) AddCity(city *City) {
	w.cities[city.Name] = city
}

func (w *World) DeleteCity(name string) {
	city, found := w.cities[name]
	if found {
		city.RemoveAllRoads()
		delete(w.cities, name)
	}
}

func (w *World) HasCity(name string) bool {
	_, ok := w.cities[name]
	return ok
}

// Parses a file into a World object.
func BuildWorld(fileName string) *World {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	world := NewWorld()

	for scanner.Scan() {
		line := scanner.Text()
		split := strings.Split(line, " ")

		// Add the new city to the graph. this assumes that there is only 1 line
		// per city in the input.
		src := NewCity(split[0])
		world.AddCity(src)

		// Process all the roads and build edges
		for i := 1; i < len(split); i++ {
			s := strings.Split(split[i], "=")
			direction := s[0]
			name := s[1]

			// If a destination city doesn't exist, create it.
			city, ok := world.FindCity(name)
			if !ok {
				city = NewCity(name)
				world.AddCity(city)
			}

			// Wire it up to the src, bi-directionally.
			src.AddRoad(city, direction)
		}
	}
	return world
}

func NewWorld() *World {
	world := &World{}
	world.cities = make(map[string]*City)
	return world
}
