package main

import (
	"reflect"
	"testing"
)

type TestParameters struct {
	name string
	args struct {
		filename string
	}
	mostExpensiveItem Record
	mostRatedItem     Record
}

func TestFindTopItems(t *testing.T) {
	tests := []TestParameters{
		testJSONReading(),
		testCSVReading(),
		testEmptyJSON(),
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1, got2 := FindTopItems(tt.args.filename)
			if !reflect.DeepEqual(got1, tt.mostExpensiveItem) {
				t.Errorf("FindTopItems() mostExpensiveItem = %v, want %v", got1, tt.mostExpensiveItem)
			}
			if !reflect.DeepEqual(got2, tt.mostRatedItem) {
				t.Errorf("FindTopItems() mostRatedItem = %v, want %v", got2, tt.mostRatedItem)
			}
		})
	}
}

func testJSONReading() TestParameters {
	newTest := TestParameters{name: "Testing JSON file processing"}
	newTest.args.filename = "testdata/db.json"
	newTest.mostExpensiveItem = Record{
		Product: "Варенье",
		Price:   200,
		Rating:  5,
	}
	newTest.mostRatedItem = Record{
		Product: "Сушки",
		Price:   100,
		Rating:  7,
	}
	return newTest
}

func testCSVReading() TestParameters {
	newTest := TestParameters{name: "Testing CSV file processing"}
	newTest.args.filename = "testdata/db.csv"
	newTest.mostExpensiveItem = Record{
		Product: "Печенье",
		Price:   3,
		Rating:  5,
	}
	newTest.mostRatedItem = Record{
		Product: "Арбуз",
		Price:   3,
		Rating:  8,
	}
	return newTest
}

func testEmptyJSON() TestParameters {
	newTest := TestParameters{name: "Testing empty JSON file"}
	newTest.args.filename = "testdata/empty.json"
	newTest.mostExpensiveItem = Record{}
	newTest.mostRatedItem = Record{}
	return newTest
}
