/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"slices"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list all entries",
	Long:  `This command lists all entries in the phone book application.`,
	Run: func(cmd *cobra.Command, args []string) {
		list()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func list() {
	entryComp := func(a, b Entry) int {
		comp := cmp.Compare(a.Surname, b.Surname)
		if comp == 0 {
			return cmp.Compare(a.Name, b.Name)
		}

		return comp
	}
	slices.SortFunc(data, entryComp)

	text, err := PrettyPrintJSONStream(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(text)
	fmt.Printf("%d records in total.\n", len(data))
}

func PrettyPrintJSONStream(data any) (string, error) {
	buffer := new(bytes.Buffer)
	encoder := json.NewEncoder(buffer)
	encoder.SetIndent("", "\t")

	err := encoder.Encode(data)
	if err != nil {
		return "", err
	}

	return buffer.String(), nil
}
