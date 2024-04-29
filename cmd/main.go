package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Alvaroalonsobabbel/csvupdate/pkg/csvupdate"
)

const (
	noFilesErr    = "You must provide two files as arguments. Use 'csvupdate -help' for help."
	updateErr     = "You must provide the fields to be updated using the -update flag. Use 'csvupdate -help' for help."
	updateFlag    = "update"
	updateHelp    = "Fields to be updated, comma separated."
	compareByErr  = "You must provide the field to compare both CSVs using the -compareby flag. Use 'csvupdate -help' for help."
	compareByFlag = "compareby"
	compareByHelp = "Field to compare both CSV files by."
	helpFlag      = "help"
	helpHelp      = "Shows this Help :)"
)

var (
	compareBy string
	update    string
)

func init() {
	flag.BoolFunc(helpFlag, helpHelp, printHelp)
	flag.StringVar(&compareBy, compareByFlag, "", compareByHelp)
	flag.StringVar(&update, updateFlag, "", updateHelp)
	flag.Parse()
}

func main() {
	if flag.NFlag() == 0 {
		err := printHelp("")
		if err != nil {
			log.Fatalf(err.Error())
		}
		return
	}

	args := flag.Args()
	if len(args) != 2 {
		log.Fatalf(noFilesErr)
	}

	if compareBy == "" {
		log.Fatal(compareByErr)
	} else if update == "" {
		log.Fatal(updateErr)
	}

	updateTool, err := csvupdate.NewUpdateTool(args[0], args[1], compareBy, update)
	if err != nil {
		log.Fatalf(err.Error())
	}

	if err := updateTool.UpdateCSV(); err != nil {
		log.Fatalf(err.Error())
	}

	fmt.Printf(
		`Updated '%s' with '%s'
Compared CSVs using the '%s' field
Updated fields: '%s'
Output file: 'out.csv'

Thanks for using CSV Update! 🎉
`, args[0], args[1], compareBy, update)
}

func printHelp(string) error {
	help := `Welcome to CSV Update. 
This tool will help you update fields in a given CSV from another CSV.

- Both CSV files must have headers.
- Headers used to compare the values by and to be updated should exist in both CSV and have the exact same name.
- Result will be generated in the 'out.csv' file.

Flags:

Arguments: 

Usage: 
  csvupdate -compareby=<field> -update=<field(s)> <outdated csv file> <updated csv file>

Examples:
  Updating RATING field based on UUID:
    csvupdate -compareby=UUID -update=RATING file_without_ratings.csv file_with_ratings.csv

  Updating RATING and COMMENTS fields based on UUID:
    csvupdate -compareby=UUID -update=RATING,COMMENTS file_outdated.csv file_with_updates.csv
`
	fmt.Println(help)
	fmt.Println("Options:")
	flag.PrintDefaults()

	os.Exit(0)
	return nil
}
