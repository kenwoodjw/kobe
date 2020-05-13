package main

import (
	"fmt"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"kobe/pkg/config"
	"kobe/pkg/inventory"
	"os"
)

var rootCmd = &cobra.Command{
	Use:   "inventory",
	Short: "A inventory provider for kobe",
	Run: func(cmd *cobra.Command, args []string) {
		host := viper.GetString("server.host")
		port := viper.GetInt("server.port")
		provider := inventory.NewKobeInventoryProvider(host, port)
		list, err := cmd.Flags().GetBool("list")
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if list {
			result, err := provider.ListHandler()
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Println(result)
			os.Exit(0)
		}
	},
}

func init() {
	rootCmd.Flags().Bool("list", false, "")
}

func main() {
	config.Init()
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}

}
