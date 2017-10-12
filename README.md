# alien-swarm

Welcome to Alienswarm!

## Usage

You can download the binary for your OS from the releases, or build from source:

```
go get github.com/eastside-eng/alienswarm
cd $GOPATH/src/github.com/eastside-eng/alienswarm
make build
```

A binary will be under `dist/` and sample input files are under `sample/`.

```
> ./dist/alienswarm -h
Alienswarm is a Go binary that runs a discrete-event simulation.

The simulation spawns a number of aliens into the world. Each iteration, the aliens move to a neighboring
city. If two or more aliens end up in the same city, they kill each other and destroy the city. The simulation will
end when all aliens are dead or the max number of iterations is exceeded. When finished, it will print the state of the
world to STDOUT.

To use alienswarm, pass an input file of the format:

	Foo north=Bar west=Baz south=Qu-ux
	Bar south=Foo west=Bee

You can use the generate command to generate random world inputs, see 'alienswarm generate -h'.

Usage:
  alienswarm <input-file> [flags]
  alienswarm [command]

Available Commands:
  generate    Generates a random NxM world to STDOUT.
  help        Help about any command

Flags:
  -h, --help                help for alienswarm
  -a, --numAliens int       The number of aliens to spawn randomly. (default 10)
  -n, --numIterations int   The number of iterations to run. (default 10000)

Use "alienswarm [command] --help" for more information about a command.
```

### Generating Worlds

I added a handy utility to generate worlds of variable size and density for
testing different implementations. You can use this by running:

```
> ./dist/alienswarm generate -h
This tool generates a random world, allowing the creator to specify the density and size of the world.

Usage:
  alienswarm generate [flags]

Flags:
  -d, --densityLimit float32   The probability of a city existing. The higher the number, the more dense the grid. (default 0.5)
  -s, --densitySeed int        A int seed to use for the generator. (default 1)
  -h, --help                   help for generate
  -l, --length int             The N-size of the grid to generate (default 10)
  -w, --width int              The M-size of the grid to generate (default 10)
```

Capture the output and plugin it back into alienswarm.

### Releases

Being new to the Golang ecosystem, I was curious how to build artifacts for other architectures and found the excellent gox utility.

I'm working on getting CircleCI automated to build new cross-arch releases on git push.

## Thought Log

The `alien-swarm` binary accepts an input file describing the cities and the roads connecting them. An additional parameter, `num-aliens`, describes the number of aliens attacking the cities. I've added an optional `seed` parameter to allow users to run the program with consistent behavior. The binary runs the discrete-event simulation up to `10,000` iterations, or until no aliens remain.

The binary emits output to STDOUT describing the remaining cities after the alien swarm has completed.

Some first-pass assumptions:

Movement
* Each alien moves a single random step. The rate of movement is fixed and they must move at least 1 city away. If an alien is at a city that has no outgoing roads, it will stay there.
* All roads are bi-directional.

Input/Output
* Implicit cities in the input will be output as explicit cities. The example input given has cities with links to them, but do not have an explicit line in the input file. These will be turned into _explicit_ cities in the output.

Spawning
* Aliens have equal probability to spawn in any location, regardless of any other alien's spawn locations. This could result in all aliens spawning at a single city and destroying each other (unlikely, given `(1/num_cities)^num_aliens`). Different distributions could be exposed as command-line parameters for further enhancement.

Destruction
* If more than 1 alien is at a city, all aliens at the city are destroyed along with the city.

Design:

The input maps directly to a cartesian plane / grid, but we can solve this more generally with a graph. There are vague invariants that each city only has at most a single road leading out in the cardinal directions, though.

My first thought is to process the file into a in-memory graph, with each node representing a city, it's status (destroyed/healthy) and a list of roads. Then, we would generate N aliens and enter an event-loop. Each iteration in the loop would iterate through all the aliens and take a random step, checking for collisions and updating the city/alien state.

This performs well, benchmarks:

```
21:43 $ make bench
go test github.com/eastside-eng/alienswarm/swarm -bench=.
goos: darwin
goarch: amd64
pkg: github.com/eastside-eng/alienswarm/swarm
BenchmarkSwarm100x100x1000-8   	   10000	    153176 ns/op
BenchmarkSwarm100x100x500-8    	   10000	    114270 ns/op
BenchmarkSwarm100x100x100-8    	   50000	     34894 ns/op
BenchmarkSwarm50x50x1000-8     	 2000000	      1083 ns/op
BenchmarkSwarm50x50x100-8      	20000000	        76.2 ns/op
BenchmarkSwarm50x50x10-8       	20000000	        75.7 ns/op
BenchmarkSwarm50x50x2-8        	20000000	       744 ns/op
BenchmarkSwarm50x50x1-8        	 3000000	       415 ns/op
ok  	github.com/eastside-eng/alienswarm/swarm	28.510s
```

Improvements I want to make:

* Make it deterministic. This involves having each alien use a different RNG that is seeded with a configurable seed.
* Refactor the Sim type into a generic harness. I'd like to put all the state into a bag and avoid having it mutated in random
places of the step function. Ideally the Step function would produce a new State in pure fashion. Short of something like persistent vectors, we can't get away from mutating the Sim state directly without losing performance or copying all the data each time we step.
* Refactor the execution model to use message passing for edge-triggered state changes. Right now, a performance bottleneck is looping through all the aliens in the Step function. We can avoid this by having Aliens emit events when they move to a city with another alien on it.

Adding channels and waitGroups slowed things down and complicated the code:

```
18:01 $ make bench
go test github.com/eastside-eng/alienswarm/swarm -bench=.
goos: darwin
goarch: amd64
pkg: github.com/eastside-eng/alienswarm/swarm
BenchmarkSwarm100x100x1000-8   	       1	4270835889 ns/op
BenchmarkSwarm100x100x500-8    	       1	2567734537 ns/op
BenchmarkSwarm100x100x100-8    	       1	1103463139 ns/op
BenchmarkSwarm50x50x1000-8     	 2000000	       622 ns/op
BenchmarkSwarm50x50x100-8      	 2000000	       602 ns/op
BenchmarkSwarm50x50x10-8       	 2000000	       583 ns/op
BenchmarkSwarm50x50x2-8        	 2000000	       857 ns/op
BenchmarkSwarm50x50x1-8        	 2000000	       705 ns/op
ok  	github.com/eastside-eng/alienswarm/swarm	20.906s
```

I'm certain we can parallelize this if we push the collision checking into the alien goroutine but I'm not going to pursue this further.

I've found that in most cases the graph converges into disjointed islands pretty quickly. Detecting if another alien is in the island will allow the sim to short circuit and I think maintaining a quad tree of all aliens would provide an efficient way to do this.
