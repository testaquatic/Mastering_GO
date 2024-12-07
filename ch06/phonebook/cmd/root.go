/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "phonebook",
	Short: "A phone book application",
	Long:  `This is a Phone Book application that uses JSON records.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := setJSONFILE()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = readJSONFile(JSONFILE)
	if err != nil && err != io.EOF {
		return
	}
	createIndex()

	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.phonebook.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

type Entry struct {
	Name       string `json:"name"`
	Surname    string `json:"surname"`
	Tel        string `json:"tel"`
	LastAccess string `json:"lastaccess"`
}

var JSONFILE = "./data.json"

type Phonebook []Entry

var data = Phonebook{}
var index map[string]int

func readJSONFile(filepath string) error {
	_, err := os.Stat(filepath)
	if err != nil {
		return err
	}

	f, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = DeSerialize(&data, f)
	if err != nil {
		return err
	}

	return nil
}

func setJSONFILE() error {
	filepath := os.Getenv("PHONEBOOK")
	if filepath != "" {
		JSONFILE = filepath
	}

	_, err := os.Stat(JSONFILE)
	if err != nil {
		fmt.Println("Creating", JSONFILE)
		f, err := os.Create(JSONFILE)
		if err != nil {
			return err
		}
		defer f.Close()
	}

	fileinfo, err := os.Stat(JSONFILE)
	if err != nil {
		return err
	}
	mode := fileinfo.Mode()
	if !mode.IsRegular() {
		return fmt.Errorf("%s not a regular file", JSONFILE)
	}

	return nil
}

func saveJSONFile(filepath string) error {
	f, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer f.Close()

	err = Serialize(&data, f)
	if err != nil {
		return err
	}

	return nil
}

func createIndex() {
	index = make(map[string]int)
	for i, k := range data {
		key := k.Tel
		index[key] = i
	}
}

func Serialize(slices any, w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(slices)
}

func DeSerialize(slice any, r io.Reader) error {
	e := json.NewDecoder(r)
	return e.Decode(slice)
}

func initS(N, S, T string) *Entry {
	if T == "" || S == "" {
		return nil
	}

	LastAccess := strconv.FormatInt(time.Now().Unix(), 10)
	return &Entry{Name: N, Surname: S, Tel: T, LastAccess: LastAccess}
}
