package swarm

import (
  "math/rand"
)

type Alien struct {
  Name string
  City *City
}

func NewAlien(name string, city *City) *Alien {
  alien := &Alien{Name: name, City: city}
  return alien
}

// Moves the Alien randomly down one of the current city's roads, if possible.
func (a *Alien) RandomWalk() {
  neighbors := a.City.GetNeighbors()
  if len(neighbors) > 0 {
    choice := rand.Intn(len(neighbors))
    a.City = neighbors[choice]
  }
}
