/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// csvCountCmd represents the csvCount command
var csvFilePath string

var csvCountCmd = &cobra.Command{
	Use:   "csvcount",
	Short: "Counts number of entries in a csv file.",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		openfile, err := os.Open(csvFilePath)
		if err != nil {
			log.Println(err)
			return
		}
		filedata, err := csv.NewReader(openfile).ReadAll()
		if err != nil {
			log.Println(err)
			return
		}
		totalRows := len(filedata)
		fmt.Println("Total #: of rows:", totalRows)
		return
	},
}

func init() {
	rootCmd.AddCommand(csvCountCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// csvCountCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// csvCountCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
