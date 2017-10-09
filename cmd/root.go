package cmd

import (
	"fmt"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	swarm "github.com/eastside-eng/alienswarm/swarm"
)

var cfgFile string
var numAliens int
var numIterations int
// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "alienswarm <input-file>",
	Short: "A discrete-event simulation for Alien Invasion.",
	Long: `Alienswarm is a Go binary that runs a discrete-event simulation.

The simulation spawns a number of aliens into the world. Each iteration, the aliens move to a neighboring
city. If two or more aliens end up in the same city, they kill each other and destroy the city. The simulation will
end when all aliens are dead or the max number of iterations is exceeded. When finished, it will print the state of the
world to STDOUT.

To use alienswarm, pass an input file of the format:

	Foo north=Bar west=Baz south=Qu-ux
	Bar south=Foo west=Bee

You can use the generate command to generate random world inputs, see 'alienswarm generate -h'.
`,
	Args: cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		sim := swarm.NewSim(args[0], numAliens, numIterations)
		sim.Run()
		sim.Graph.PrintGraph()
	},
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	RootCmd.Flags().IntVarP(&numAliens, "numAliens", "a", 10, "The number of aliens to spawn randomly.")
	RootCmd.Flags().IntVarP(&numIterations, "numIterations", "n", 10000, "The number of iterations to run.")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".alienswarm" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".alienswarm")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
