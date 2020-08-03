package cmd

import (
	"fmt"
	"log"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const appname = "stine"
const description = "STiNE CLI"
const longDescription = ""

// defaultCfgFile is the default config file name without extension
const defaultCfgFile = "." + appname
const defaultCfgFileType = "yaml"
const envPrefix = appname

// cfgFile contains the config file path if set by a CLI flag
var cfgFile string

// printVersion is true when version flag is set
var printVersion bool

// verbose is a flag to enable verbose output
var verbose bool

// semester is a flag used to specify a semester for: exams, examresults
var semester string

// output is a flag used to specify an output file for: scheduler export
var output string

// rootCmd represents the base command when called without any sub-commands
var rootCmd = &cobra.Command{
	Use:   appname,
	Short: description,
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			versionCmd.Run(cmd, args)
			return
		}

		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	log.SetFlags(log.Ltime)

	// Define flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "Config file (default is "+defaultCfgFile+"."+defaultCfgFileType+" in program dir, CWD or $HOME)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "Username")
	rootCmd.PersistentFlags().Bool("nocache", true, "Disable session cache")
	rootCmd.PersistentFlags().BoolVar(&verbose, "verbose", false, "Provide a more verbose output")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "Show version and exit")

	// Bind flags to config values
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("nocache", rootCmd.PersistentFlags().Lookup("nocache"))

	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.SetDefault("nocache", false)
}

// initConfig reads in config file and environment variables if set.
func initConfig() {
	if printVersion {
		// skip reading config when printing version
		return
	}

	if cfgFile != "" {
		// Use config file from the flag
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Default config name
		viper.SetConfigName(defaultCfgFile)
		// Default config type
		viper.SetConfigType(defaultCfgFileType)

		if ex, err := os.Executable(); err == nil {
			// Search for config in directory of executable
			viper.AddConfigPath(ex)
		}

		// Search for config in CWD
		viper.AddConfigPath(".")
		// Search for config in home dir
		viper.AddConfigPath(home)
	}

	// Read in environment variables that match
	viper.AutomaticEnv()
	viper.SetEnvPrefix(envPrefix)

	// If a config file is found, read it in
	if err := viper.ReadInConfig(); err == nil && verbose {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
