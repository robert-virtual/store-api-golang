package main

import "fmt"

func createProduct(product product) (*product, error) {
	_, err := db.Exec(
		"INSERT INTO products(name,description,price,quantity,userId) values(?,?,?,?,?)",
		product.Name,
		product.Description,
		product.Price,
		product.Quantity,
		product.UserId,
	)
	insertImages(product)
	if err != nil {
		return nil, fmt.Errorf("createProduct Error %v", err)
	}
	return &product, nil
}
func findProducts() ([]product, error) { // retorna una lista de usuarios y un error que muede ser nulo
	var products []product
	rows, err := db.Query("select * from products")
	if err != nil {
		return nil, fmt.Errorf("findProducts Error %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var product product
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.UserId, &product.CreatedAt); err != nil {
			return nil, fmt.Errorf("findProducts Error %v", err)
		}
		if product.UserId != nil {
			product.User, err = findUserById(*product.UserId)
			fmt.Println(err)
		}
		images, err := findImages(*product.Id)
		if err == nil {
			product.Images = images
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findProducts Error %v", err)
	}
	return products, nil
}

func findProductId(product product) (*string, error) {
	row := db.QueryRow(
		"select id from products where name = ? and description = ? and userId = ? and price = ? and quantity = ?",
		product.Name,
		product.Description,
		product.UserId,
		product.Price,
		product.Quantity,
	)
	if err := row.Scan(&product.Id); err != nil {

		return nil, fmt.Errorf("findUserByEmail Error %v", err)
	}
	return product.Id, nil
}
