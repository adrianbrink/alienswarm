# alien-swarm
A coding example in Go with a CI pipeline and docker container setup.

## $?!

The `alien-swarm` binary accepts an input file describing the cities and the roads connecting them. An additional parameter, `num-aliens`, describes the number of aliens attacking the cities. I've added an optional `seed` parameter to allow users to run the program with consistent behavior. The binary runs the discrete-event simulation up to `10,000` iterations, or until no aliens remain.

The binary emits output to STDOUT describing the remaining cities after the alien swarm has completed.

Some first-pass assumptions:

Movement
* Each alien moves a single random step. The rate of movement is fixed and they must move atleast 1 city away. If an alien is at a city that has no outgoing roads, it will stay there. All roads are bi-directional.

Input/Output
* Implicit cities in the input will be output as explicit cities. The example input given has cities with links to them, but do not have an explicit line in the input file. These will be turned into _explicit_ cities in the output.

Spawning
* Aliens have equal probability to spawn in any location, regardless of any other alien's spawn locations. This could result in all aliens spawning at a single city and destroying each other (unlikely, given `(1/num_cities)^num_aliens`). Different distributions could be exposed as command-line parameters for further enhancement.

Destruction
* If more than 1 alien is at a city, all aliens at the city are destroyed along with the city.

Design:

The input maps directly to a cartesian plane / grid, but we can solve this more generally with a graph. There are vague invariants that each city only has at most a single road leading out in the cardinal directions, though.

My first thought is to process the file into a in-memory graph, with each node representing a city, it's status (destroyed/healthy) and a list of roads. Then, we would generate N aliens and enter an event-loop. Each iteration in the loop would iterate through all the aliens and take a random step, checking for collisions and updating the city/alien state.
