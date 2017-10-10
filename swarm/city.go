package swarm

import "sync"

// Using a City struct to model the city graph. Each city will
// have bi-directional edges in the cardinal directions. An additional
// boolean `Destroyed` represents the state of the city during simulation.
type City struct {
	Name  string
	North *City
	East  *City
	South *City
	West  *City

	// In the concurrent version, each alien will add/remove themselves from the
	// visitor set during their step routine. We use a simple mutex to sync. Since
	// the read will happen in the update step, we're guaranteed reads will happen
	// after all writes.
	visitorLock sync.Mutex
	visitors    []*Alien

	NumVisitors int
}

func (src *City) String() string {
	str := src.Name
	if src.North != nil {
		str = str + " north=" + src.North.Name
	}
	if src.East != nil {
		str = str + " east=" + src.East.Name
	}
	if src.South != nil {
		str = str + " south=" + src.South.Name
	}
	if src.West != nil {
		str = str + " west=" + src.West.Name
	}
	return str
}

// Bi-directionally adds a edge to the city graph.
func (src *City) AddRoad(dst *City, direction string) {
	switch direction {
	case "north":
		src.North = dst
		dst.South = src
	case "east":
		src.East = dst
		dst.West = src
	case "south":
		src.South = dst
		dst.North = src
	case "west":
		src.West = dst
		dst.East = src
	}
}

func (src *City) RemoveRoad(direction string) {
	switch direction {
	case "north":
		if src.North != nil {
			dst := src.North
			src.North = nil
			dst.South = nil
		}
	case "east":
		if src.East != nil {
			dst := src.East
			src.East = nil
			dst.West = nil
		}
	case "south":
		if src.South != nil {
			dst := src.South
			src.South = nil
			dst.North = nil
		}
	case "west":
		if src.West != nil {
			dst := src.West
			src.West = nil
			dst.East = nil
		}
	}
}

func (src *City) RemoveAllRoads() {
	src.RemoveRoad("north")
	src.RemoveRoad("east")
	src.RemoveRoad("south")
	src.RemoveRoad("west")
}

func (src *City) GetNeighbors() []*City {
	neighbors := make([]*City, 0)
	if src.North != nil {
		neighbors = append(neighbors, src.North)
	}
	if src.East != nil {
		neighbors = append(neighbors, src.East)
	}
	if src.South != nil {
		neighbors = append(neighbors, src.South)
	}
	if src.West != nil {
		neighbors = append(neighbors, src.West)
	}
	return neighbors
}

func (src *City) AddVisitor(alien *Alien) {
	src.visitorLock.Lock()
	src.visitors = append(src.visitors, alien)
	src.NumVisitors = len(src.visitors)
	src.visitorLock.Unlock()
}

func (src *City) RemoveVisitor(alien *Alien) {
	src.visitorLock.Lock()
	for i, a := range src.visitors {
		if a == alien {
			src.visitors = append(src.visitors[:i], src.visitors[i+1:]...)
			break
		}
	}
	src.NumVisitors = len(src.visitors)
	src.visitorLock.Unlock()
}

// GetVisitors does not lock because this should only be invoked on the Update step of the Sim.
// All visitor mutations happen in the Step step.
func (src *City) GetVisitors() []*Alien {
	return src.visitors
}

func NewCity(name string) *City {
	city := &City{Name: name}
	return city
}
