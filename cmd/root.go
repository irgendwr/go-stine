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
const description = "STiNE CLI tool"
const longDescription = ""

// defaultCfgFile is the default config file name without extention
const defaultCfgFile = "." + appname
const defaultCfgFileType = "yaml"
const envPrefix = appname

// nolint: gochecknoglobals
var (
	version = "dev"
	commit  = ""
	date    = ""
	builtBy = ""
)

// cfgFile contains the config file path if set by a CLI flag
var cfgFile string

// printVersion is true when version flag is set
var printVersion bool

// semester is a flag used to specify a semester for: exams, examresults
var semester string

// output is a flag used to specify an output file for: sheduler export
var output string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   appname,
	Short: description,
	Long:  longDescription,
	Run: func(cmd *cobra.Command, args []string) {
		if printVersion {
			fmt.Println(description)
			fmt.Println(buildVersion(version, commit, date))
			return
		}

		cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	log.SetFlags(log.Ltime)

	// Define flags
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "c", "", "config file (default is "+defaultCfgFile+"."+defaultCfgFileType+" in program dir, CWD or $HOME)")
	rootCmd.PersistentFlags().StringP("username", "u", "", "username")
	rootCmd.PersistentFlags().BoolP("nocache", "n", true, "disable session cache")
	rootCmd.Flags().BoolVarP(&printVersion, "version", "v", false, "show version and exit")

	// Bind flags to config values
	viper.BindPFlag("username", rootCmd.PersistentFlags().Lookup("username"))
	viper.BindPFlag("nocache", rootCmd.PersistentFlags().Lookup("nocache"))

	viper.SetDefault("username", "")
	viper.SetDefault("password", "")
	viper.SetDefault("nocache", false)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if printVersion {
		// skip reading config when printVersionis set
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
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func buildVersion(version, commit, date string) string {
	var result = fmt.Sprintf("version: %s", version)
	if commit != "" {
		result = fmt.Sprintf("%s\ncommit: %s", result, commit)
	}
	if date != "" {
		result = fmt.Sprintf("%s\nbuilt at: %s", result, date)
	}
	if builtBy != "" {
		result = fmt.Sprintf("%s\nbuilt by: %s", result, builtBy)
	}
	return result
}
