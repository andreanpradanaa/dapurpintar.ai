package model

type Difficulty string

const (
	DifficultyEasy   Difficulty = "easy"
	DifficultyMedium Difficulty = "medium"
	DifficultyHard   Difficulty = "hard"
)

type DietaryTag string

const (
	DietaryVegetarian DietaryTag = "vegetarian"
	DietaryVegan      DietaryTag = "vegan"
	DietaryHalal      DietaryTag = "halal"
	DietaryGlutenFree DietaryTag = "gluten-free"
	DietarySpicy      DietaryTag = "spicy"
	DietaryLowCarb    DietaryTag = "low-carb"
)

type Category string

const (
	CategoryProtein   Category = "protein"
	CategoryVegetable Category = "vegetable"
	CategorySpice     Category = "spice"
	CategoryGrain     Category = "grain"
	CategoryDairy     Category = "dairy"
	CategorySauce     Category = "sauce"
	CategoryOther     Category = "other"
)

type RecipeIngredient struct {
	Name     string  `json:"name"`
	NameID   string  `json:"nameId"`
	Amount   string  `json:"amount"`
	Category Category `json:"category"`
	Optional bool    `json:"optional,omitempty"`
}

type RecipeStep struct {
	Order       int    `json:"order"`
	Text        string `json:"text"`
	TextID      string `json:"textId"`
	DurationSec *int   `json:"durationSec,omitempty"`
	Tip         string `json:"tip,omitempty"`
}

type Nutrition struct {
	Calories int `json:"calories"`
	Protein  int `json:"protein"`
	Carbs    int `json:"carbs"`
	Fat      int `json:"fat"`
	Fiber    int `json:"fiber"`
}

type Recipe struct {
	ID           string             `json:"id"`
	Slug         string             `json:"slug"`
	Title        string             `json:"title"`
	TitleID      string             `json:"titleId"`
	Description  string             `json:"description"`
	DescriptionID string             `json:"descriptionId"`
	Image        string             `json:"image,omitempty"`
	Gradient     []string           `json:"gradient"`
	Cuisine      string             `json:"cuisine"`
	Difficulty   Difficulty         `json:"difficulty"`
	PrepTime     int                `json:"prepTime"`
	CookTime     int                `json:"cookTime"`
	Servings     int                `json:"servings"`
	Ingredients  []RecipeIngredient `json:"ingredients"`
	Steps        []RecipeStep       `json:"steps"`
	Nutrition    Nutrition          `json:"nutrition"`
	Tags         []string           `json:"tags"`
	Dietary      []DietaryTag       `json:"dietary"`
	Rating       float64            `json:"rating"`
	Reviews      int                `json:"reviews"`
	CreatedAt    string             `json:"createdAt"`
}
