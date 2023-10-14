/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	//"errors"
	"fmt"
	"os/exec"

	"github.com/spf13/cobra"
)

// rcalculateCmd represents the rcalculate command
var rcalculateCmd = &cobra.Command{
	Use:   "rcalculate",
	Short: "Calculate Summary Statistics with R",
	Long:  "R will read the input.csv, calculate the max, min, and mean values, then summarize these summary statistics in an output.txt file.",

	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("Rscript", "./runHouses.R").Output()

		//print error if it occurs
		if err != nil {
			fmt.Println(err)
		}

		//use out variable to avoid error
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(rcalculateCmd)
}
