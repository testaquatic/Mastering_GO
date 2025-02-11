/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/spf13/cobra"
)

// getidCmd represents the getid command
var getidCmd = &cobra.Command{
	Use:   "getid",
	Short: "Returns User ID, given a username",
	Long:  `This command returns the User ID of user, given their username`,
	Run: func(cmd *cobra.Command, args []string) {
		endpoint := "/getid"
		user := User{Username: username, Password: password}

		var u2 User
		err := json.Unmarshal([]byte(data), &u2)
		if err != nil {
			fmt.Println("Unmarshal:", err)
			return
		}

		if u2.Username == "" {
			fmt.Println("Empty username!")
			return
		}

		buf := new(bytes.Buffer)
		err = user.ToJSON(buf)
		if err != nil {
			fmt.Println("ToJSON:", err)
			return
		}

		URL := SERVER + PORT + endpoint + "/" + u2.Username
		req, err := http.NewRequest(http.MethodGet, URL, buf)
		if err != nil {
			fmt.Println("NewRequest:", err)
			return
		}
		req.Header.Add("Content-Type", "application/json")

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

		var Returned = User{}
		err = Returned.FromJSON(resp.Body)
		if err != nil {
			fmt.Println("SliceFromJSON:", err)
			return
		}

		fmt.Println("User", Returned.Username, "has ID:", Returned.ID)
	},
}

func init() {
	rootCmd.AddCommand(getidCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getidCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getidCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
