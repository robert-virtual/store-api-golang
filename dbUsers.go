package main

import "fmt"

func createUser(user user) (*user, error) {
	hashedPassword := "luego lo hacemos"
	_, err := db.Exec(
		"INSERT INTO users(name,email,password,image) values(?,?,?,?)",
		user.Name,
		user.Email,
		hashedPassword,
		user.Image,
	)
	if err != nil {
		return nil, fmt.Errorf("createuser Error %v", err)
	}
	return &user, nil
}
func findUsers() ([]user, error) { // retorna una lista de usuarios y un error que muede ser nulo
	var users []user
	rows, err := db.Query("select * from users")
	if err != nil {
		return nil, fmt.Errorf("findUsers Error %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var user user
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.Image); err != nil {
			return nil, fmt.Errorf("findUsers Error %v", err)
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("findUsers Error %v", err)
	}
	return users, nil
}