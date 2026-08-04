package llm

import (
	"strings"
	"testing"
)

func minimalValidRecipe() string {
	return `{"title":"Garlic Chicken and Egg Stir-Fry","titleId":"Garlic Chicken and Egg Stir-Fry","description":"A quick Indonesian-style stir-fry of tender chicken and scrambled eggs with plenty of garlic.","descriptionId":"Tumis ayam dan telur ala Indonesia dengan banyak bawang putih.","cuisine":"Indonesian","difficulty":"easy","prepTime":10,"cookTime":15,"servings":2,"ingredients":[{"name":"Chicken breast","nameId":"Dada ayam","amount":"200g","category":"protein","optional":false},{"name":"Eggs","nameId":"Telur","amount":"2","category":"protein","optional":false},{"name":"Garlic","nameId":"Bawang putih","amount":"4 cloves","category":"spice","optional":false}],"steps":[{"order":1,"text":"Slice chicken into thin strips.","textId":"Iris ayam tipis-tipis.","durationSec":180,"tip":""},{"order":2,"text":"Heat oil and fry garlic until fragrant.","textId":"Panaskan minyak dan tumis bawang putih hingga harum.","durationSec":60,"tip":""},{"order":3,"text":"Add chicken, cook until no longer pink.","textId":"Masukkan ayam, masak hingga berubah warna.","durationSec":300,"tip":"Do not overcook."}],"nutrition":{"calories":320,"protein":38,"carbs":5,"fat":18,"fiber":1},"tags":["stir-fry","chicken","quick"],"dietary":["halal","low-carb"]}`
}

func minimalValidRecipeNoLastField() string {
	return `{"title":"Garlic Chicken and Egg Stir-Fry","titleId":"Garlic Chicken and Egg Stir-Fry","description":"A quick Indonesian-style stir-fry","descriptionId":"Tumis ayam dan telur","cuisine":"Indonesian","difficulty":"easy","prepTime":10,"cookTime":15,"servings":2,"ingredients":[{"name":"Chicken breast","nameId":"Dada ayam","amount":"200g","category":"protein","optional":false}],"steps":[{"order":1,"text":"Slice chicken.","textId":"Iris ayam.","durationSec":180}],"nutrition":{"calories":320,"protein":38,"carbs":5,"fat":18,"fiber":1}}`
}

// --- Repair Strategy Unit Tests ---

func TestIdentityRepair(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantOK  bool
	}{
		{"valid json", `{"a":1}`, `{"a":1}`, true},
		{"empty string", ``, ``, true},
		{"truncated missing brace", `{"a":1`, `{"a":1`, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := identityRepair(tt.input)
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestCloseTruncatedStringRepair(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    string
		wantOK  bool
	}{
		{
			name:   "mid-string truncation",
			input:  `{"description":"A quick`,
			want:   `{"description":"A quick"`,
			wantOK: true,
		},
		{
			name:   "mid-string trailing space",
			input:  `{"description":"A quick  `,
			want:   `{"description":"A quick"`,
			wantOK: true,
		},
		{
			name:   "already closed string",
			input:  `{"description":"done"`,
			want:   "",
			wantOK: false,
		},
		{
			name:   "ends with brace",
			input:  `{"a":1}`,
			want:   "",
			wantOK: false,
		},
		{
			name:   "ends with bracket",
			input:  `{"a":[1,2]`,
			want:   "",
			wantOK: false,
		},
		{
			name:   "empty input",
			input:  "",
			want:   "",
			wantOK: false,
		},
		{
			name:   "whitespace only",
			input:  "   \n",
			want:   "",
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := closeTruncatedStringRepair(tt.input)
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

func TestStripLastFieldRepair(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		want   string
		wantOK bool
	}{
		{
			name:   "valid field separator",
			input:  `{"title":"Food","description":"Yum","broken"`,
			want:   `{"title":"Food","description":"Yum"}`,
			wantOK: true,
		},
		{
			name:   "no field to strip",
			input:  `{"title":"Food"}`,
			want:   "",
			wantOK: false,
		},
		{
			name:   "empty",
			input:  ``,
			want:   "",
			wantOK: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, ok := stripLastFieldRepair(tt.input)
			if ok != tt.wantOK {
				t.Errorf("ok = %v, want %v", ok, tt.wantOK)
			}
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// --- tryCloseJSON ---

func TestTryCloseJSON(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"already closed", `{"a":1}`, `{"a":1}`},
		{"missing one", `{"a":1`, `{"a":1}`},
		{"missing two", `{"outer":{"inner":"val"`, `{"outer":{"inner":"val"}}`},
		{"nested in array", `{"steps":[{"order":1}`, `{"steps":[{"order":1}]}`},
		{"empty", ``, ``},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tryCloseJSON(tt.input)
			if got != tt.want {
				t.Errorf("got %q, want %q", got, tt.want)
			}
		})
	}
}

// --- parseAndHydrate End-to-End ---

func TestParseAndHydrate_WellFormed(t *testing.T) {
	r, err := parseAndHydrate(minimalValidRecipe())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Title != "Garlic Chicken and Egg Stir-Fry" {
		t.Errorf("title = %q", r.Title)
	}
	if r.Difficulty != "easy" {
		t.Errorf("difficulty = %q", r.Difficulty)
	}
	if len(r.Ingredients) != 3 {
		t.Errorf("ingredients = %d, want 3", len(r.Ingredients))
	}
	if len(r.Steps) != 3 {
		t.Errorf("steps = %d, want 3", len(r.Steps))
	}
	if r.Nutrition.Calories != 320 {
		t.Errorf("calories = %d", r.Nutrition.Calories)
	}
	if len(r.Tags) != 3 {
		t.Errorf("tags = %d", len(r.Tags))
	}
}

func TestParseAndHydrate_MissingCloseBrace(t *testing.T) {
	raw := minimalValidRecipe()
	// Drop the final }
	truncated := raw[:len(raw)-1]
	r, err := parseAndHydrate(truncated)
	if err != nil {
		t.Fatalf("unexpected error on missing close brace: %v", err)
	}
	if r.Title != "Garlic Chicken and Egg Stir-Fry" {
		t.Errorf("title = %q", r.Title)
	}
}

func TestParseAndHydrate_MidStringTruncation(t *testing.T) {
	raw := minimalValidRecipe()
	// Cut mid-string inside the description, plus drop closing brace
	idx := strings.Index(raw, `"A quick`) + len(`"A quick `)
	if idx <= 0 {
		t.Fatal("could not find truncation point")
	}
	truncated := raw[:idx+5] // "A quick" + "Indon"
	r, err := parseAndHydrate(truncated)
	if err != nil {
		t.Fatalf("unexpected error on mid-string truncation: %v", err)
	}
	if r.Title != "Garlic Chicken and Egg Stir-Fry" {
		t.Errorf("title = %q", r.Title)
	}
	if !strings.HasPrefix(r.Description, "A quick Indon") {
		t.Errorf("description = %q, want prefix 'A quick Indon'", r.Description)
	}
}

func TestParseAndHydrate_TruncatedWithCodeFence(t *testing.T) {
	raw := "```json\n" + minimalValidRecipeNoLastField() + "\n```"
	r, err := parseAndHydrate(raw)
	if err != nil {
		t.Fatalf("unexpected error on code-fenced json: %v", err)
	}
	if r.Title != "Garlic Chicken and Egg Stir-Fry" {
		t.Errorf("title = %q", r.Title)
	}
}

func TestParseAndHydrate_NoStepsOrIngredients_filledByHydrate(t *testing.T) {
	r, err := parseAndHydrate(`{"title":"Only name"}`)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if r.Title != "Only name" {
		t.Errorf("title = %q", r.Title)
	}
	if len(r.Ingredients) == 0 {
		t.Error("ingredients should not be empty after hydrate")
	}
	if len(r.Steps) == 0 {
		t.Error("steps should not be empty after hydrate")
	}
}

func TestParseAndHydrate_GarbageInput(t *testing.T) {
	_, err := parseAndHydrate("not json at all")
	if err == nil {
		t.Fatal("expected error on garbage input")
	}
}

func TestParseAndHydrate_DeeplyNestedTruncation(t *testing.T) {
	raw := `{"title":"Test","description":"Testing` // no ingredients, steps, nutrition — mid-string on description
	r, err := parseAndHydrate(raw)
	if err != nil {
		t.Fatalf("unexpected error on mid-string description truncation: %v", err)
	}
	if r.Title != "Test" {
		t.Errorf("title = %q", r.Title)
	}
	// ingredients and steps should be filled by hydrate defaults
	if len(r.Ingredients) == 0 {
		t.Error("ingredients should not be empty")
	}
}
