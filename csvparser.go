package main

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func readCSV(filename string, aggregator *Aggregator) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed opening a file: %s", err)
	}
	defer f.Close()

	separator := defineSeparator(f)

	stream := bufio.NewReaderSize(f, 16*1024*1024) // reading through a 16 MB buffer
	reader := csv.NewReader(stream)
	// Replace Comma to another separator
	reader.Comma = separator

	// Read header
	header, err := reader.Read()
	if err != nil {
		log.Fatalf("File %s has wrong format: %s", filename, err)
	}
	// Check header
	if len(header) != 3 || header[0] != "Product" || header[1] != "Price" || header[2] != "Rating" {
		log.Fatalf("The header in the file %s has wrong format. It must be: 'Product,Price,Rating'", filename)
	}

	// While the array contains values
	for {
		line, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatal(err)
		}
		// Convert string values to numbers
		price, err := strconv.ParseInt(line[1], 10, 64)
		if err != nil {
			log.Fatalf("Failed parsing a price %s: %s", line[1], err)
		}
		rating, err := strconv.ParseInt(line[2], 10, 64)
		if err != nil {
			log.Fatalf("Failed parsing a rating %s: %s", line[2], err)
		}
		// Compose Record
		rec := Record{
			Product: line[0],
			Price:   money(price),
			Rating:  int(rating),
		}
		// Execute aggregation
		aggregator.ExamineRecord(rec)
	}

}

func defineSeparator(f *os.File) rune {
	scanner := bufio.NewScanner(f)
	// Reset a file reading
	defer f.Seek(0, io.SeekStart)
	// Read the first string
	if scanner.Scan() {
		line := scanner.Text()
		// Try to identify the separator used (e.g. it may be wether "," or ";")
		if strings.HasPrefix(line, "Product") {
			separator := line[len("Product")]
			return rune(separator)
		}
	}
	return ',' // default separator
}
