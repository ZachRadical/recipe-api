package repos

import (
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"

)

func CreateRecipe(recipe models.Recipe) (models.Recipe, error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO recipes (name, servings, user_id)
	VALUES ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(sqlStatement, recipe.Name, recipe.Servings, recipe.UserID).Scan(&recipe.ID)
	if err != nil {
		return recipe, err
	}
	return recipe, nil
}

func UpdateRecipe(recipe models.Recipe, id int) (models.Recipe, error) {
	db := database.GetDB()
	sqlStatement := `UPDATE recipes SET name = $2, servings = $3 WHERE id = $1
	RETURNING id, name, servings, user_id`
	
	err := db.QueryRow(sqlStatement, id, recipe.Name, recipe.Servings).Scan(&recipe.ID, &recipe.Name, &recipe.Servings, &recipe.UserID)
	if err != nil {
		return models.Recipe{}, err
	}
	recipe.ID = id
	return recipe, nil
}

func GetRecipe(recipe models.Recipe, id int) (models.Recipe, error){
	db := database.GetDB()
	sqlStatement := `SELECT id, name, servings, user_id FROM recipes WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&recipe.ID, &recipe.Name, &recipe.Servings, &recipe.UserID)
	if err != nil {
		return models.Recipe{}, err
	}
	return recipe, nil
}

// func GetAllRecipesByUser(user_id int) ([]models.Recipe, error){
// 	db := database.GetDB()
// 	sqlStatement := `SELECT name, servings FROM recipes WHERE user_id = $1`

// 	err := db.QueryRow(sqlStatement, user_id).Scan(&user_id)
// 	if err != nil {
// 		return []models.Recipe{}, err
// 	}
// 	return []models.Recipe{}, nil
// }