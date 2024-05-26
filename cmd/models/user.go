package models

import "time"

type User struct {
	ID			int				`json:"id"`
	Username	string			`json:"username"`
	Email		string			`json:"email"`
	Password	string			`json:"password"`
	CreatedAt	time.Time		`json:"created_at"`
	UpdatedAt	time.Time		`json:"updated_at"`
}

type Recipe struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Servings	float64			`json:"servings"`
	UserID		int				`json:"user_id"`
}

type Component struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Directions	string			`json:"directions"`
	RecipeID	int				`json:"recipe_id"`
}

type Ingredient struct {
	ID			int				`json:"id"`
	Name		string			`json:"name"`
	Quantity	float64			`json:"quantity"`
	Unit		string			`json:"unit"`
	ComponentID	int				`json:"component_id"`
}
