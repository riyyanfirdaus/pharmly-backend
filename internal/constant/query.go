package constant

const (
	QCreateUser = `
		INSERT INTO
			users (username, full_name, email, password, role, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5,$6, $7)
		RETURNING id
	`

	QGetByEmail = `
		SELECT
			id, username, full_name, email, password, role, status, created_at, updated_at, deleted_at
		FROM
			users
		WHERE
			email = $1
	`

	QGetByID = `
		SELECT
			id, username, full_name, email, password, role, status, created_at, updated_at, deleted_at
		FROM
			users
		WHERE
			id = $1
	`

	QGetAllUsers = `
		SELECT
			id, username, full_name, email, password, role, status, created_at, updated_at, deleted_at
		FROM
			users
		ORDER BY
			updated_at
		DESC
		LIMIT
			$1
		OFFSET
			$2
	`

	QCountUserQuery = `
		SELECT
			COUNT(*)
		FROM
			users
	`

	QGetAllCategories = `
		SELECT
			id, name, description, parent_category_id, created_at, updated_at, deleted_at
		FROM
			categories
		ORDER BY
			updated_at
		DESC
		LIMIT
			$1
		OFFSET
			$2
	`

	QCountCategoryQuery = `
		SELECT 
			COUNT(*)
		FROM
			categories
	`

	QCreateProduct = `
		INSERT INTO
			products (name, category_id, generic_name, description, price, stock, unit, expiration_date, barcode, supplier_id, min_stock, is_active, created_at, updated_at)
		VALUES
			($1, $2, $3, $4, $5,$6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING id
	`

	QGetAllProducts = `
		SELECT
			id, name, category_id, generic_name, description, price, stock, unit, expiration_date, barcode, supplier_id, min_stock, is_active, created_at, updated_at, deleted_at
		FROM
			products
		ORDER BY
			updated_at
		DESC
		LIMIT
			$1
		OFFSET
			$2
	`

	QCountProductQuery = `
		SELECT
			COUNT(*)
		FROM
			products
	`

	QGetAllSuppliers = `
		SELECT
			id, name, contact_person, phone, address, email, created_at, updated_at, deleted_at
		FROM
			suppliers
		ORDER BY
			updated_at
		DESC
		LIMIT
			$1
		OFFSET
			$2
	`

	QCountSupplierQuery = `
		SELECT COUNT(*)
		FROM
			suppliers
	`
)
