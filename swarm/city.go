package swarm

// Using a City struct to model the city graph. Each city will
// have bi-directional edges in the cardinal directions. An additional
// boolean `Destroyed` represents the state of the city during simulation.
type City struct {
	Name  string
	North *City
	East  *City
	South *City
	West  *City
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

func NewCity(name string) *City {
	city := &City{Name: name}
	return city
}
