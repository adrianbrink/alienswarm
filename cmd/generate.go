package cmd

import (
	"github.com/spf13/cobra"

	swarm "github.com/eastside-eng/alienswarm/swarm"
)

var n int
var m int
var densityLimit float32
var densitySeed int

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generates a random NxM world to STDOUT.",
	Long: `Super useful for testing and benchmarking alienswarm.
	This tool generates a random world, allowing the creator to specify the density and size of the world.`,
	Run: func(cmd *cobra.Command, args []string) {
		world := swarm.Generate(n, m, densityLimit, densitySeed)
		swarm.PrintWorld(n,m, world)
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)
	generateCmd.Flags().IntVarP(&n, "length", "l", 10, "The N-size of the grid to generate")
	generateCmd.Flags().IntVarP(&m, "width", "w", 10, "The M-size of the grid to generate")
	generateCmd.Flags().Float32VarP(&densityLimit, "densityLimit", "d", .50,
		 "A float32 representing the probability of a city existing. The higher the number, the more dense the grid.")
	generateCmd.Flags().IntVarP(&densitySeed, "densitySeed", "s", 1, "A int seed to use for the generator.")
}
