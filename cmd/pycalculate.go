/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	// "errors"
	"fmt"
	"os/exec"

	// "strings"

	"github.com/spf13/cobra"
)

// pycalculateCmd represents the pycalculate command
var pycalculateCmd = &cobra.Command{
	Use:   "pycalculate",
	Short: "Calculate Summary Statistics with Python",
	Long:  "Python will read the input.csv, calculate the max, min, and mean values, then summarize these summary statistics in an output.txt file.",

	Run: func(cmd *cobra.Command, args []string) {
		out, err := exec.Command("python", "./runHouses.py").Output()

		//print errors if they occur
		if err != nil {
			fmt.Println(err)
		}

		//use out variable to avoid error
		fmt.Println(string(out))
	},
}

func init() {
	rootCmd.AddCommand(pycalculateCmd)
}
