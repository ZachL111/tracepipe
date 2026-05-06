package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 69, Capacity: 91, Latency: 21, Risk: 20, Weight: 7}, wantScore: 61, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 95, Capacity: 73, Latency: 25, Risk: 9, Weight: 8}, wantScore: 166, wantDecision: "accept"},
		{name: "case_3", signal: Signal{Demand: 97, Capacity: 90, Latency: 22, Risk: 17, Weight: 8}, wantScore: 137, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
