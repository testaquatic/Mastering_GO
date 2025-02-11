/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// loggedCmd represents the logged command
var loggedCmd = &cobra.Command{
	Use:   "logged",
	Short: "List add logged in users",
	Long:  `This command shows all logged in users`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/logged"
		user := User{Username: username, Password: password}

		buf := new(bytes.Buffer)
		err := user.ToJSON(buf)
		if err != nil {
			fmt.Println("JSON:", err)
			return
		}

		req, err := http.NewRequest(http.MethodGet, SERVER+PORT+endpoint, buf)
		if err != nil {
			fmt.Println("NewRequest:", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		c := &http.Client{
			Timeout: 15 * time.Second,
		}

		resp, err := c.Do(req)
		if err != nil {
			fmt.Println("Do:", err)
			return
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println(resp)
			return
		}

		var users = []User{}
		err = SliceFromJSON(&users, resp.Body)
		if err != nil {
			fmt.Println("SliceFromJSON:", err)
			return

		}
		data, err := PrettyJSON(users)
		if err != nil {
			fmt.Println("PrettyJSON:", err)
			return
		}
		fmt.Println(string(data))
	},
}

func init() {
	rootCmd.AddCommand(loggedCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// loggedCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// loggedCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
