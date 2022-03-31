package main

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
)

func ReadJSON(filename string, aggregator *Aggregator) {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("Failed opening a file: %s", err)
	}
	defer f.Close()

	stream := bufio.NewReaderSize(f, 16*1024*1024) // reading through a 16 MB buffer
	decoder := json.NewDecoder(stream)

	// Read the open bracket
	t, err := decoder.Token()
	if err == io.EOF {
		log.Printf("%s doesn't have JSON data", filename)
		return
	} else if t != json.Delim('[') {
		log.Fatalf("File %s has wrong format. The first token must be '[', instead of '%s'", filename, t)
	}

	// While the array contains values
	for decoder.More() {
		var rec Record
		err := decoder.Decode(&rec)
		if err != nil {
			log.Fatalf("Failed decoding the record: %s", err)
		}
		aggregator.ExamineRecord(rec)
	}

	// Read the close bracket
	t, err = decoder.Token()
	if err == io.EOF || t != json.Delim(']') {
		log.Fatalf("File %s has wrong format. The last token must be ']', instead of '%s'", filename, t)
	}
}
