package repos

import (
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"

)

func CreateIngredient(ingredient models.Ingredient) (models.Ingredient, error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO ingredients (name, quantity, unit, component_id)
	VALUES ($1, $2, $3, $4) RETURNING id`
	
	err := db.QueryRow(sqlStatement, ingredient.Name, ingredient.Quantity, ingredient.Unit, ingredient.ComponentID).Scan(&ingredient.ID)
	if err != nil {
		return ingredient, err
	}
	return ingredient, nil
}

func UpdateIngredient(ingredient models.Ingredient, id int) (models.Ingredient, error) {
	db := database.GetDB()
	sqlStatement := `UPDATE ingredients SET name = $2, quantity = $3, unit = $4 WHERE id = $1`
	
	err := db.QueryRow(sqlStatement, id, ingredient.Name, ingredient.Quantity, ingredient.Unit).Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit)
	if err != nil {
		return models.Ingredient{}, err
	}
	ingredient.ID = id
	return ingredient, nil
}

func GetIngredient(ingredient models.Ingredient, id int) (models.Ingredient, error){
	db := database.GetDB()
	sqlStatement := `SELECT id, name, quantity, unit, component_id FROM ingredients WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&ingredient.ID, &ingredient.Name, &ingredient.Quantity, &ingredient.Unit, &ingredient.ComponentID)
	if err != nil {
		return models.Ingredient{}, err
	}
	return ingredient, nil
}

// func GetAllIngredientsByComponent(component_id int) ([]models.Ingredient, error){
// 	db := database.GetDB()
// 	sqlStatement := `SELECT name, quantity, unit FROM ingredients WHERE component_id = $1`

// 	err := db.QueryRow(sqlStatement, component_id).Scan(&component_id)
// 	if err != nil {
// 		return []models.Ingredient{}, err
// 	}
// 	return []models.Ingredient{}, nil
// }