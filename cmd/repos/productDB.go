package repos

import (
	"nomcom-api/cmd/database"
	"nomcom-api/cmd/models"

)

func CreateProduct(product models.Product) (models.Product, error) {
	db := database.GetDB()
	sqlStatement := `INSERT INTO product (name, size, price)
	VALUES ($1, $2, $3) RETURNING id`
	
	err := db.QueryRow(sqlStatement, product.Name, product.Size, product.Price).Scan(&product.ID)
	if err != nil {
		return product, err
	}
	return product, nil
}

func UpdateProduct(product models.Product, id int) (models.Product, error) {
	db := database.GetDB()
	sqlStatement := `UPDATE product SET name = $2, size = $3, price = $4 WHERE id = $1
	RETURNING id, name, size, price`
	
	err := db.QueryRow(sqlStatement, id, product.Name, product.Size, product.Price).Scan(&product.ID, &product.Name, &product.Size, &product.Price)
	if err != nil {
		return models.Product{}, err
	}
	product.ID = id
	return product, nil
}

func GetProduct(product models.Product, id int) (models.Product, error){
	db := database.GetDB()
	sqlStatement := `SELECT name, size, price FROM product WHERE id = $1`

	err := db.QueryRow(sqlStatement, id).Scan(&product.ID, &product.Name, &product.Size, &product.Price)
	if err != nil {
		return models.Product{}, err
	}
	return product, nil
}

// func GetAllProducts(products []models.Product) ([]models.Product, error){
// 	db := database.GetDB()
// 	sqlStatement := `SELECT name, size, price FROM product`

// 	err := db.QueryRow(sqlStatement).Scan(&product.Name, product.Size, product.Price)
// 	if err != nil {
// 		return []models.Product{}, err
// 	}
// 	return []models.Product{}, nil
// }