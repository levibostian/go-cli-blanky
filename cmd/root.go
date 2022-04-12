package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	cliConfig "github.com/levibostian/Purslane/cliconfig"
	"github.com/levibostian/Purslane/ui"
)

var cfgFile string
var debug bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "purslane",
	Short: "Scalable, flexible, affordable, and safe way to perform periodic tasks.",
	Long: `Scalable, flexible, affordable, and safe way to perform periodic tasks.
	
Purslane is quite simple, really. Each time Purslane is run, it will:
1. Create a new cloud server and optionally attach disk storage to it. (Purslane works with [DigitalOcean](https://www.digitalocean.com/) at this time but may include other cloud providers in the future)
2. Run a Docker container of your choice in the newly created cloud server. 
3. When the Docker container exits, Purslane will destroy all of the resources it created such as the created server and disk storage.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) {
	// },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		ui.HandleError(err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here, will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is ~/.purslane)")

	rootCmd.PersistentFlags().BoolVar(&debug, "debug", false, "Show debug statements. Used for debugging program for bug reports and development. (default false)")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		ui.Message("Using config file: " + viper.ConfigFileUsed())
	}

	cliConfig.SetCliConfig(debug)
}
