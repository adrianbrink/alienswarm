package swarm

import (
  "os"
  "fmt"
  "bufio"
  "log"
  "strings"
  "math/rand"
  "strconv"
)

// Working solution.
// TODO: Implement a linked list for easy deletion of destroyed cities and
// dead aliens. This will allow us to avoid checking each city and alien
// in the event loop.

// Using a City struct to model the city graph. Each city will
// have bi-directional edges in the cardinal directions. An additional
// boolean `Destroyed` represents the state of the city during simulation.
type City struct {
  Name  string
  North *City
  East  *City
  South *City
  West  *City
  Destroyed bool
}

func (src *City) String() string {
  str := src.Name
  if src.Destroyed {
    str = str + " DESTROYED"
  }
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

func (src *City) GetNeighbors() []*City {
  neighbors := make([]*City, 0)
  if src.North != nil && !src.North.Destroyed {
    neighbors = append(neighbors, src.North)
  }
  if src.East != nil && !src.East.Destroyed {
    neighbors = append(neighbors, src.East)
  }
  if src.South != nil && !src.South.Destroyed {
    neighbors = append(neighbors, src.South)
  }
  if src.West != nil && !src.West.Destroyed {
    neighbors = append(neighbors, src.West)
  }
  return neighbors
}

func NewCity(name string) *City {
  city := &City{Name: name, Destroyed: false}
  return city
}

type Graph struct {
  cities map[string]*City
}

func NewGraph() *Graph {
  graph := &Graph{}
  graph.cities = make(map[string]*City)
  return graph
}

func (g *Graph) PrintGraph() {
  for _, v := range g.cities {
    fmt.Println(v)
  }
}

func (g *Graph) FindCity(name string) (*City, bool) {
  city, found := g.cities[name]
  return city, found
}

func (g *Graph) RandomCity() *City {
  choice := rand.Intn(len(g.cities))
  count := 0
  // The order here is undefined for `range` -- we randomly choose an index
  // but if the order happens to have some relationship with rand.Intn then
  // this could be come deterministic.
  for _,v := range g.cities {
    if count == choice {
      return v
    }
    count++
  }
  return nil
}

func (g *Graph) AddCity(city *City) {
  g.cities[city.Name] = city
}

func (g *Graph) DeleteCity(key string) {
  delete(g.cities, key)
}

func (g *Graph) HasCity(name string) bool {
  _, ok := g.cities[name]
  return ok
}

func BuildGraph(fileName string) *Graph {
  file, err := os.Open(fileName)
  if err != nil {
      log.Fatal(err)
  }
  defer file.Close()

  scanner := bufio.NewScanner(file)
  graph := NewGraph()

  for scanner.Scan() {
    line := scanner.Text()
    split := strings.Split(line, " ")

    // Add the new city to the graph. this assumes that there is only 1 line
    // per city in the input.
    src := NewCity(split[0])
    graph.AddCity(src)

    // Process all the roads and build edges
    for i := 1; i < len(split); i++ {
      s := strings.Split(split[i], "=")
      direction := s[0]
      name := s[1]

      // If a destination city doesn't exist, create it.
      city, ok := graph.FindCity(name)
      if !ok {
        city = NewCity(name)
        graph.AddCity(city)
      }

      // Wire it up to the src, bi-directionally.
      src.AddRoad(city, direction)
    }
  }
  return graph
}

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

type Sim struct {
  Graph *Graph
  Aliens []*Alien
  MaxIterations int
  CurrentIteration int
}

func (sim *Sim) Step() {
  // Move all aliens
  for _,a := range sim.Aliens {
    a.RandomWalk()
  }

  // Build up a map of cities -> aliens
  occupancy := make(map[string][]*Alien)
  for _,a := range sim.Aliens {
    name := a.City.Name
    aliens, found := occupancy[name]
    if !found {
      aliens = make([]*Alien, 0)
    }
    occupancy[name] = append(aliens, a)
  }

  newAliens := make([]*Alien, 0)

  // For the occupancy map, if len(aliens) > 1, destroy everything.
  for _,v := range occupancy {
    numAliens := len(v)
    if numAliens > 1 {
      names := make([]string, numAliens)
      c := v[0].City
      c.Destroyed = true
      for i,a := range v {
        names[i] = a.Name
      }
      // Builds up a string for each alien, handling the edge case at the end.
      killList := strings.Join(names[:numAliens-1], ", ")
      killList += " and " + v[numAliens-1].Name
      // fmt.Printf("%s was destroyed by aliens %s!\n", c.Name, killList)

    } else {
      newAliens = append(newAliens, v...)
    }
  }

  sim.Aliens = newAliens
}

func (sim *Sim) Run() {
  for sim.CurrentIteration <= sim.MaxIterations && len(sim.Aliens) != 0 {
    sim.Step()
    sim.CurrentIteration++
  }
}

func NewSim(cityFile string, numAliens int, maxIterations int) *Sim {
  s := &Sim{MaxIterations: maxIterations}
  s.Graph = BuildGraph(cityFile)
  s.Aliens = make([]*Alien, numAliens)
  for i := 0; i < numAliens; i++ {
    s.Aliens[i] = NewAlien(strconv.Itoa(i), s.Graph.RandomCity())
  }
  s.CurrentIteration = 0
  return s
}
