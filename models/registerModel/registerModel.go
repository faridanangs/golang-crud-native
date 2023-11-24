package registermodel

import (
	"crud_go_native/config"
	"crud_go_native/entities"
	"crud_go_native/helpers"
)

func Create(user entities.Register) bool {
	query, err := config.DB.Exec(`
		insert into user(name, password) values(?, ?)
	`, user.Name, user.Password)
	helpers.FuncError(err, "error exec db at create user sign up model")

	result, _ := query.RowsAffected()

	return result > 0
}

func UserAll() []entities.Register {
	rows, err := config.DB.Query(`
		select name, password from user
	`)
	helpers.FuncError(err, "error query db at UserAll model")
	defer rows.Close()

	var users []entities.Register

	for rows.Next() {
		var user entities.Register
		err := rows.Scan(&user.Name, &user.Password)
		helpers.FuncError(err, "error scan row at UserAll model")
		users = append(users, user)
	}

	return users
}
