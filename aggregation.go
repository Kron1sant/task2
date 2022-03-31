package main

// Aggregator allows to aggregate the most expensive and most rated product.
// If several products have the same value, then take the first one
type Aggregator struct {
	MostExpensiveItem Record
	MostRatedItem     Record
}

func newAggregator() *Aggregator {
	return &Aggregator{}
}

// ExamineRecord compares the current values with the new record
func (a *Aggregator) ExamineRecord(rec Record) {
	if rec.Price > a.MostExpensiveItem.Price {
		a.MostExpensiveItem = rec
	}
	if rec.Rating > a.MostRatedItem.Rating {
		a.MostRatedItem = rec
	}
}
