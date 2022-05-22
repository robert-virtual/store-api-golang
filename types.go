package main

// nombremos las propiedades en ingles mejor
// los nombres de las propiedades deben estar en mayusculkas para poder ser accedidos desde otros acerchivos
type user struct {
	Id       *string   `json:"id"` // apuntador porque este campo puede no estar por ejemplo cuando recivamos un nuevo producto a insertar antes de insertar no tendra id
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"password"`
	Image    *string   `json:"image"`    //url
	Products []product `json:"products"` // lista de productos
}
type product struct {
	Id          *string  `json:"id"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Price       float64  `json:"price"`
	Quantity    int64    `json:"quantity"`
	UserId      *string  `json:"userId"`
	User        user     `json:"user"`
	Images      *[]image `json:"images"` // *puntero = es nullable(puede ser nulo)
}
type image struct {
	Id        *string `json:"id"` // al convertirse a json la propiedad Id sera renombrada a id
	Url       string  `json:"url"`
	ProductId *string `json:"productId"`
}
