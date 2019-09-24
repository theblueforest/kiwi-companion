package commands

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/theblueforest/kiwi-companion/helpers"
)

var configFile string

var rootCmd = &cobra.Command{
	Use:   "kiwi",
	Short: "Automate your code projects with a unique CLI",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&configFile, "config", "", "config file (default is $HOME/.kiwi-companion.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func initConfig() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		root := helpers.ConfigsGetRootPath()

		// Directories
		if _, err := os.Stat(root); os.IsNotExist(err) {
			os.Mkdir(root, 0750)
			os.Mkdir(helpers.ConfigsGetKubernetesPath(root), 0750)
		}

		// Configs
		viper.AddConfigPath(root)
		viper.SetConfigName(".config")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
