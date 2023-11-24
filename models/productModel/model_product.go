package productmodel

import (
	"crud_go_native/config"
	"crud_go_native/entities"
	"crud_go_native/helpers"
)

func GetAll() []entities.Product {
	rows, err := config.DB.Query(`
		SELECT 
			products.id,
			products.name,
			categories.name as category_name,
			products.stock,
			products.description,
			products.created_at,
			products.updated_at
		FROM products
		JOIN categories ON products.category_id = categories.id
	`)
	helpers.FuncError(err, "query db error at get all products model")
	defer rows.Close()

	var products []entities.Product

	for rows.Next() {
		var product entities.Product
		err := rows.Scan(
			&product.Id,
			&product.Name,
			&product.Category.Name,
			&product.Stock,
			&product.Description,
			&product.CreatedAt,
			&product.UpdatedAt,
		)
		helpers.FuncError(err, "error rows scan at get all product model")
		products = append(products, product)
	}

	return products
}

func Detail(id int) entities.Product {
	row := config.DB.QueryRow(`
	SELECT 
		products.id,
		products.name,
		categories.name as category_name,
		products.stock,
		products.description,
		products.created_at,
		products.updated_at
	FROM products
	JOIN categories ON products.category_id = categories.id
	where products.id = ?
	`, id)

	var product entities.Product

	err := row.Scan(
		&product.Id,
		&product.Name,
		&product.Category.Name,
		&product.Stock,
		&product.Description,
		&product.CreatedAt,
		&product.UpdatedAt,
	)
	helpers.FuncError(err, "error row scan at detail product model")

	return product

}

func Create(product entities.Product) bool {
	result, err := config.DB.Exec(`
	insert into products(
		name, category_id, stock, description, created_at, updated_at
	) values (?, ?, ?, ?, ?, ?)
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		product.CreatedAt,
		product.UpdatedAt,
	)
	helpers.FuncError(err, "exect db error at create product model")

	productID, _ := result.LastInsertId()

	return productID > 0
}

func Update(id int, product entities.Product) bool {
	query, err := config.DB.Exec(`
		update products set
		name = ?,
		category_id = ?,
		stock = ?,
		description = ?,
		updated_at = ?
		where id = ?
	`,
		product.Name,
		product.Category.Id,
		product.Stock,
		product.Description,
		product.UpdatedAt,
		id,
	)
	helpers.FuncError(err, "error exec db at update product model")

	result, _ := query.RowsAffected()
	return result > 0
}

func Delete(id int) bool {
	query, err := config.DB.Exec("delete from products where products.id = ?", id)
	helpers.FuncError(err, "error exec db at delete products model")

	result, _ := query.RowsAffected()

	return result > 0
}
