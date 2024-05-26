package repos

import (
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"

)

func CreateUser(user models.User) (models.User,error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO users (username, email, password)
	VALUES ($1, $2, $3) RETURNING id`

	err := db.QueryRow(sqlStatement, user.Username, user.Email, user.Password).Scan(&user.ID)

	if err != nil {
		return user, err
	}
	return user, nil
}