package repos

import (
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"
)

func CreateComponent(component models.Component) (models.Component, error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO components (name, directions, recipe_id)
	VALUES ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(sqlStatement, component.Name, component.Directions, component.RecipeID).Scan(&component.ID)
	if err != nil {
		return component, err
	}
	
	return component, nil
}

func UpdateComponent(component models.Component, id int) (models.Component, error) {
	db := database.GetDB()
	sqlStatement := `UPDATE components SET name = $2, directions = $3 WHERE id = $1
	RETURNING id, name, directions, recipe_id`
	
	err := db.QueryRow(sqlStatement, id, component.Name, component.Directions).Scan(&component.ID, &component.Name, &component.Directions, &component.RecipeID)
	if err != nil {
		return models.Component{}, err
	}
	component.ID = id
	return component, nil
}

func GetComponent(component models.Component, id int) (models.Component, error){
	db := database.GetDB()
	sqlStatement := `SELECT id, name, directions, recipe_id FROM components WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&component.ID, &component.Name, &component.Directions, &component.RecipeID)
	if err != nil {
		return models.Component{}, err
	}
	return component, nil
}

// func GetAllComponentsByRecipe(recipe_id int) ([]models.Component, error){
// 	db := database.GetDB()
// 	sqlStatement := `SELECT name, directions FROM components WHERE recipe_id = $1`

// 	err := db.QueryRow(sqlStatement, recipe_id).Scan(&recipe_id)
// 	if err != nil {
// 		return []models.Component{}, err
// 	}
// 	return []models.Component{}, nil
// }