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
		if err := rows.Scan(&product.Id, &product.Name, &product.Description, &product.Price, &product.Quantity, &product.UserId); err != nil {
			return nil, fmt.Errorf("findProducts Error %v", err)
		}
		if product.UserId != nil {
			product.User, err = findUserById(*product.UserId)
			fmt.Println(err)
		}
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findProducts Error %v", err)
	}
	return products, nil
}
