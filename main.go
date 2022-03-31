package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
)

type TypeOfDataFile uint

const (
	CSV TypeOfDataFile = iota + 1
	JSON
)

type money int

type Record struct {
	Product string `json:"product"`
	Price   money  `json:"price"`
	Rating  int    `json:"rating"`
}

func main() {
	// The filename is specified in the startup arguments
	filename := getFilenameFromCLI()
	if _, err := os.Stat(filename); err != nil {
		log.Fatalf("The file does not exist: %s", err)
	}

	// Looking for the desired values
	mostExpensiveItem, mostRatedItem := findTopItems(filename)

	// Printing the result
	fmt.Printf("The most expensive product %s with a price of %d\n", mostExpensiveItem.Product, mostExpensiveItem.Price)
	fmt.Printf("The most rated product %s with a rating of %d\n", mostRatedItem.Product, mostRatedItem.Rating)
}

func getFilenameFromCLI() string {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n \t %s <filename(.json|.csv)>\n\n", os.Args[0])
		flag.PrintDefaults()
	}
	flag.Parse()
	args := flag.Args()
	// It's allowed to specify only one file name
	if len(args) != 1 {
		flag.Usage()
		os.Exit(1)
	}

	return args[0]
}

func findTopItems(filename string) (Record, Record) {

	// Use an aggregator that implements the logic of
	// looking for the most expensive and the most rated items
	aggregator := newAggregator()

	// Choose a parser depending on the file type
	// and then give the filename and the aggregator to the parser
	switch getTypeOfFile(filename) {
	case CSV:
		readCSV(filename, aggregator)
	case JSON:
		readJSON(filename, aggregator)
	}

	return aggregator.MostExpensiveItem, aggregator.MostRatedItem
}

func getTypeOfFile(filename string) TypeOfDataFile {
	ext := strings.ToLower(path.Ext(filename))
	switch ext {
	case ".csv":
		return CSV
	case ".json":
		return JSON
	default:
		log.Fatalf("Unknown file extension: %s", ext)
	}
	return 0
}
