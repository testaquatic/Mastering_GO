/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// timeCmd represents the time command
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "Get the time for  the RESTful server",
	Long:  `This command mainly exists for making sure that the server works`,
	Run: func(cmd *cobra.Command, args []string) {
		req, err := http.NewRequest(http.MethodGet, SERVER+PORT+"/time", nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}

		if resp == nil || (resp.StatusCode == http.StatusNotFound) {
			fmt.Println(resp)
			return
		}
		defer resp.Body.Close()

		data, _ := io.ReadAll(resp.Body)
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(timeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// timeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// timeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
