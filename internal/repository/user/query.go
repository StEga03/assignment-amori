package user

const (
	// Used for reference table.
	tableDefUsers = "users"

	// List of queries for users.
	querySelectUsersClause = `SELECT
									id,
									first_name,
									last_name,
									birth_date,
									gender,
									gender_interest,
									phone_number,
									relationship_status,
									relationship_goal,
									created_at,
									updated_at
								FROM ` + tableDefUsers

	querySelectUserById = querySelectUsersClause + ` WHERE id = $1`
)
