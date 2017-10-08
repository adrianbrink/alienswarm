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
var cityFile string
var numAliens int
var numIterations int

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "alienswarm",
	Short: "A brief description of your application",
	Long: `...`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Swarm called")
		sim := swarm.NewSim(cityFile, numAliens, numIterations)
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

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.alienswarm.yaml)")

	// Pull inputs from flags or env/config via Viper.
	RootCmd.Flags().StringVarP(&cityFile, "inputFile", "f", "", "The input file describe cities and their roads.")
	viper.BindPFlag("inputFile", RootCmd.PersistentFlags().Lookup("inputFile"))

	RootCmd.Flags().IntVarP(&numAliens, "numAliens", "a", 10, "The number of aliens to spawn randomly.")
	viper.BindPFlag("numAliens", RootCmd.PersistentFlags().Lookup("numAliens"))

	RootCmd.Flags().IntVarP(&numIterations, "numIterations", "n", 10000, "The number of iterations to run.")
	viper.BindPFlag("numIterations", RootCmd.PersistentFlags().Lookup("numIterations"))
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
