/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

var InputFileName string
var OutputFileName string

// gocalculateCmd represents the gocalculate command
var gocalculateCmd = &cobra.Command{
	Use:   "gocalculate",
	Short: "Calculate Summary Statistics with Go",
	Long:  "Go will read the input.csv, calculate the max, min, and mean values, then summarize these summary statistics in an output.txt file.",

	Run: func(cmd *cobra.Command, args []string) {
		if InputFileName == "" || OutputFileName == "" {
			fmt.Println("Please specify input and output file names")
			return
		}

		//Open the input CSV file
		InputFileName, err := os.Open(InputFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer InputFileName.Close()

		//Create CSV reader
		reader := csv.NewReader(bufio.NewReader(InputFileName))

		//Read the CSV data
		records, err := reader.ReadAll()
		if err != nil {
			log.Fatal(err)
		}

		//Calculate min, max and mean for each column

		//Create empty containers
		var columnSums []float64
		var columnMin []float64
		var columnMax []float64

		//Add values to empty containers
		for i := 0; i < 7; i++ {
			columnSums = append(columnSums, 0)
			columnMin = append(columnMin, float64(9999999))
			columnMax = append(columnMax, float64(-9999999))
		}

		//iterate through each column and calculate each statistic
		for _, record := range records[1:] {
			for i, value := range record {
				num, err := strconv.ParseFloat(value, 64)
				if err != nil {
					log.Fatal(err)
				}
				columnSums[i] += num
				if num < columnMin[i] {
					columnMin[i] = num
				}
				if num > columnMax[i] {
					columnMax[i] = num
				}
			}
		}

		//Calculate the Mean
		var columnMeans []float64
		for i := 0; i < 7; i++ {
			columnMeans = append(columnMeans, columnSums[i]/float64(len(records)))
		}

		//Create the output file
		OutputFileName, err := os.Create(OutputFileName)
		if err != nil {
			log.Fatal(err)
		}
		defer OutputFileName.Close()

		//Write the summary statistics to the output file
		outputWriter := bufio.NewWriter(OutputFileName)
		for i := 0; i < 7; i++ {
			fmt.Fprint(outputWriter, "Column", " ", i+1, ":", records[0][i], "\n")
			fmt.Fprint(outputWriter, "Minimum:", " ", columnMin[i], "\n")
			fmt.Fprint(outputWriter, "Maximum:", " ", columnMax[i], "\n")
			fmt.Fprint(outputWriter, "Mean:", " ", columnMeans[i], "\n")
		}

		//Flush and close the output file
		outputWriter.Flush()
	},
}

func init() {
	rootCmd.AddCommand(gocalculateCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// gocalculateCmd.PersistentFlags().String("foo", "", "A help for foo")
	rootCmd.PersistentFlags().StringVarP(&InputFileName, "input", "i", "", "Input CSV file")
	rootCmd.PersistentFlags().StringVarP(&OutputFileName, "output", "o", "", "Output plain text file")
	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// gocalculateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
