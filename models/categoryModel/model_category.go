package categorymodel

import (
	"crud_go_native/config"
	"crud_go_native/entities"
	"crud_go_native/helpers"
)

func GetAll() []entities.Category {
	rows, err := config.DB.Query("select * from categories")
	helpers.FuncError(err, "query db error at get all category model")
	defer rows.Close()

	var categories []entities.Category

	for rows.Next() {
		var category entities.Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
		helpers.FuncError(err, "scan error at get all category")
		categories = append(categories, category)
	}
	return categories
}

func Create(category entities.Category) bool {
	result, err := config.DB.Exec(`
	Insert into categories(name, created_at, updated_at) 
	value (?, ?, ?)`,
		category.Name, category.CreatedAt, category.UpdatedAt)
	helpers.FuncError(err, "exect db error at create category model")

	categoryID, err := result.LastInsertId()
	helpers.FuncError(err, "lastIndexId db error at create category model")

	return categoryID > 0
}

func Detail(id int) entities.Category {
	rows := config.DB.QueryRow("select * from categories where id = ?", id)

	var category entities.Category

	err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.UpdatedAt)
	helpers.FuncError(err, "scan db error at detail category model")

	return category
}

func Update(id int, category entities.Category) bool {
	query, err := config.DB.Exec(`Update categories set name = ?, updated_at = ? where id = ?`, category.Name, category.UpdatedAt, id)
	helpers.FuncError(err, "exec db error at update category model")

	result, err := query.RowsAffected()
	helpers.FuncError(err, "rowsAffected db error at update category model")

	return result > 0
}

func Delete(id int) {
	_, err := config.DB.Exec("delete from categories where id = ?", id)
	helpers.FuncError(err, "Exec db error at delete category model")
}
