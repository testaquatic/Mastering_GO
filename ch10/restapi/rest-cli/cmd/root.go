/*
Copyright © 2025 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/json"
	"io"
	"os"

	"github.com/spf13/cobra"
)

var SERVER string
var PORT string
var data string
var username string
var password string

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	LastLogin int64  `json:"lastlogin"`
	Admin     int    `json:"admin"`
	Active    int    `json:"active"`
}

const (
	empty = ""
	tap   = "\t"
)

func (p *User) FromJSON(r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(p)
}

func (p *User) ToJSON(w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(p)
}

func SliceFromJSON(slice interface{}, r io.Reader) error {
	e := json.NewDecoder(r)

	return e.Decode(slice)
}

func SliceToJSON(slice interface{}, w io.Writer) error {
	e := json.NewEncoder(w)

	return e.Encode(slice)
}

func PrettyJSON(data interface{}) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent(empty, tap)
	err := encoder.Encode(data)
	if err != nil {
		return empty, err
	}

	return buffer.String(), nil
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "rest-cli",
	Short: "A REST API client",
	Long: `A Client for a RESTful server.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.rest-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "username", "The username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "admin", "The password")
	rootCmd.PersistentFlags().StringVarP(&data, "data", "d", "{}", "JSON Record")
	rootCmd.PersistentFlags().StringVarP(&SERVER, "server", "s", "http://localhost", "The server")
	rootCmd.PersistentFlags().StringVarP(&PORT, "port", "P", ":1234", "Port of RESTful Server")
}
