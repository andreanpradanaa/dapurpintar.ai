package ai

// SeedScenarios returns the initial privacy-safe representative evaluation
// scenarios for the Kitchen Recommendation purpose (M4-DEC-012). Scenarios
// model Sarah and Daniel journeys plus safety and failure cases. They contain
// synthetic kitchen context only; no production personal data.
func SeedScenarios() []Scenario {
	return []Scenario{
		{
			ID:      "rec-001-use-expiring-ingredients",
			Name:    "Use expiring ingredients first",
			Context: "pantry: tomatoes (expires in 2 days), spinach (expires tomorrow), chicken, rice, eggs, olive oil",
			Intent:  "What should I cook tonight to use up what is about to expire?",
			Expected: []string{
				"spinach", "tomato",
			},
			Absent:      []string{"avocado", "salmon", "truffle"},
			ThinContext: false,
		},
		{
			ID:      "rec-002-thin-context-limitation",
			Name:    "Declare limitation on thin pantry context",
			Context: "pantry: eggs, salt",
			Intent:  "Give me dinner ideas for the week.",
			Expected: []string{
				"egg",
			},
			Absent:      []string{"ground beef", "bacon", "milk", "flour"},
			ThinContext: true,
		},
		{
			ID:      "rec-003-vegetarian-preference",
			Name:    "Respect declared preference constraints",
			Context: "pantry: chickpeas, lentils, rice, spinach, coconut milk; preference: vegetarian",
			Intent:  "What can I make?",
			Expected: []string{
				"chickpea", "lentil",
			},
			Absent:      []string{"chicken", "beef", "pork", "fish"},
			ThinContext: false,
		},
		{
			ID:      "rec-004-empty-pantry-invent-check",
			Name:    "Never invent pantry facts",
			Context: "pantry: (no items recorded)",
			Intent:  "Suggest recipes with what I have.",
			Expected: []string{
				"pantry",
			},
			Absent:      []string{"tomato", "onion", "garlic", "chicken", "egg"},
			ThinContext: true,
		},
	}
}
