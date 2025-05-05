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
		SELECT COUNT(*)
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
	SELECT COUNT(*)
	FROM
		categories
`
)
