package swarm

import (
	"fmt"
	"math/rand"
)

// Stolen from Moby/Docker source for generating docker names.
var (
	names = [...]string{
    "albattani",
		"allen",
		"almeida",
		"agnesi",
		"archimedes",
		"ardinghelli",
		"aryabhata",
		"austin",
		"babbage",
		"banach",
		"bardeen",
		"bartik",
		"bassi",
		"beaver",
		"bell",
		"benz",
		"bhabha",
		"bhaskara",
		"blackwell",
		"bohr",
		"booth",
		"borg",
		"bose",
		"boyd",
		"brahmagupta",
		"brattain",
		"brown",
		"carson",
		"chandrasekhar",
		"shannon",
		"clarke",
		"colden",
		"cori",
		"cray",
		"curran",
		"curie",
		"darwin",
		"davinci",
		"dijkstra",
		"dubinsky",
		"easley",
		"edison",
		"einstein",
		"elion",
		"engelbart",
		"euclid",
		"euler",
		"fermat",
		"fermi",
		"feynman",
		"franklin",
		"galileo",
		"gates",
		"goldberg",
		"goldstine",
		"goldwasser",
		"golick",
		"goodall",
		"haibt",
		"hamilton",
		"hawking",
		"heisenberg",
		"hermann",
		"heyrovsky",
		"hodgkin",
		"hoover",
		"hopper",
		"hugle",
		"hypatia",
		"jackson",
		"jang",
		"jennings",
		"jepsen",
		"johnson",
		"joliot",
		"jones",
		"kalam",
		"kare",
		"keller",
		"kepler",
		"khorana",
		"kilby",
		"kirch",
		"knuth",
		"kowalevski",
		"lalande",
		"lamarr",
		"lamport",
		"leakey",
		"leavitt",
		"lewin",
		"lichterman",
		"liskov",
		"lovelace",
		"lumiere",
		"mahavira",
		"mayer",
		"mccarthy",
		"mcclintock",
		"mclean",
		"mcnulty",
		"meitner",
		"meninsky",
		"mestorf",
		"minsky",
		"mirzakhani",
		"morse",
		"murdock",
		"neumann",
		"newton",
		"nightingale",
		"nobel",
		"noether",
		"northcutt",
		"noyce",
		"panini",
		"pare",
		"pasteur",
		"payne",
		"perlman",
		"pike",
		"poincare",
		"poitras",
		"ptolemy",
		"raman",
		"ramanujan",
		"ride",
		"montalcini",
		"ritchie",
		"roentgen",
		"rosalind",
		"saha",
		"sammet",
		"shaw",
		"shirley",
		"shockley",
		"sinoussi",
		"snyder",
		"spence",
		"stallman",
		"stonebraker",
		"swanson",
		"swartz",
		"swirles",
		"tesla",
		"thompson",
		"torvalds",
		"turing",
		"varahamihira",
		"visvesvaraya",
		"volhard",
		"wescoff",
		"wiles",
		"williams",
		"wilson",
		"wing",
		"wozniak",
		"wright",
		"yalow",
		"yonath",
	}
)

func PrintWorld(n int, m int, world []string) {
  for i := 0; i < n; i++ {
    for j := 0; j < m; j++ {
      idx := i * n + j
			if world[idx] != "" {
	      out := world[idx]
	      if j != 0 && world[idx-1] != "" {
	        out += " west=" + world[idx-1] // west
	      }
	      if j != n - 1 && world[idx+1] != "" {
	        out += " east=" + world[idx+1] // east
	      }
	      if i != 0 && world[idx-m] != "" {
	        out += " north=" + world[idx-m] // north
	      }
	      if i != n - 1 && world[idx+m] != "" {
	        out += " south=" + world[idx+m] // south
	      }
	      fmt.Println(out)
			}
    }
  }
}

func Generate(n int, m int, densityLimit float32, densitySeed int) ([]string) {
  r := rand.New(rand.NewSource(int64(densitySeed)))
  world := make([]string, n*m)
  // generate an NxM grid, randomly choose coordinates to turn into cities.
  for i := range world {
    roll := r.Float32()
    if roll < densityLimit {
      nameIdx := i % len(names)
      count := (i / len(names))
			name := names[nameIdx]
			switch count {
			case 0: // do nothing
			case 1: name = "new_" + name
			case 2: name = "san_" + name
			case 3: name = "uss_" + name
			case 4: name = "moonbase_" + name
			default: name = fmt.Sprintf("%s_%d", name, count)
			}
      world[i] = name
    } else {
      world[i] = ""
    }
  }
  return world
}
