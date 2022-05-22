package main

// nombremos las propiedades en ingles mejor
type user struct {
	id       *string // apuntador porque este campo puede no estar por ejemplo cuando recivamos un nuevo producto a insertar antes de insertar no tendra id
	name     string
	email    string
	password string
	image    string    //url
	products []product // lista de productos
}
type product struct {
	id         *string
	name       string
	desciption string
	price      float64
	quantity   int64
	userId     *string
	user       user
	images     *[]image // *puntero = es nullable(puede ser nulo)
}
type image struct {
	id        *string
	url       string
	productId *string
}
