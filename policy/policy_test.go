package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 69, Capacity: 91, Latency: 21, Risk: 20, Weight: 7}
	if got := Score(signal); got != 61 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 95, Capacity: 73, Latency: 25, Risk: 9, Weight: 8}
	if got := Score(signal); got != 166 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "accept" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 97, Capacity: 90, Latency: 22, Risk: 17, Weight: 8}
	if got := Score(signal); got != 137 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
