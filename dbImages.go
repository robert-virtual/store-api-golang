package main

import "fmt"

func findImages(productId string) ([]image, error) { // retorna una lista de usuarios y un error que muede ser nulo
	var images []image
	rows, err := db.Query("select * from images where productId = ?", productId)
	if err != nil {
		return nil, fmt.Errorf("findImages Error %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var image image
		if err := rows.Scan(&image.Id, &image.Url, &image.ProductId); err != nil {
			return nil, fmt.Errorf("findImages Error %v", err)
		}
		images = append(images, image)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findImages Error %v", err)
	}
	return images, nil
}
func insertImages(product product) {

	if product.Images == nil {
		return
	}
	id, err := findProductId(product)
	if err != nil {
		return
	}
	for _, image := range product.Images {
		image.ProductId = id
		createImage(image)
	}
}
func createImage(image image) (*image, error) {
	_, err := db.Exec(
		"INSERT INTO images(url,productId) values(?,?)",
		image.Url,
		image.ProductId,
	)
	if err != nil {
		return nil, fmt.Errorf("createImage Error %v", err)
	}
	return &image, nil
}
