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