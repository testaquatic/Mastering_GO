/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search for a telephone number",
	Long:  `This command searches for a given phone number..`,
	Run: func(cmd *cobra.Command, args []string) {
		SERVER := viper.GetString("server")
		PORT := viper.GetString("port")

		number, _ := cmd.Flags().GetString("tel")
		if number == "" {
			fmt.Println("Number is empty!")
			return
		}

		URL := "http://" + SERVER + ":" + PORT + "/search/" + number

		data, err := http.Get(URL)
		if err != nil {
			fmt.Println(err)
			return
		}

		if data.StatusCode != http.StatusOK {
			fmt.Println("Status code:", data.StatusCode)
			return
		}

		responseData, err := io.ReadAll(data.Body)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(responseData))
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)
	searchCmd.Flags().StringP("tel", "t", "", "Telephone number to search")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
