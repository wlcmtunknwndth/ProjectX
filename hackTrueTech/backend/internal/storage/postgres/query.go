package postgres

const (
	// AUTH
	getPassword  = "SELECT password FROM auth WHERE username = $1"
	registerUser = "INSERT INTO auth(username, password, gender, age) VALUES($1, $2, $3, $4)"
	isAdmin      = "SELECT isadmin FROM auth WHERE username = $1"
	deleteUser   = "DELETE FROM auth WHERE username = $1"

	//Event
	getEvent    = "SELECT * FROM events WHERE id = $1"
	createEvent = `INSERT INTO events(
							price,
							restrictions,
							date,
							city,
                   			address,
							name,
							img_path,
							description
                   			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id
	`
	changeImgPath = "UPDATE events SET img_path=$1"
	patchEvent    = `UPDATE events SET price = $1,
											restrictions = $2,
											date = $3,
											city = $4,
											address = $5,
											name = $6,
											description = $7
									WHERE id = $8
											`

	deleteEvent        = "DELETE FROM events WHERE id = $1"
	getEventsByFeature = "SELECT * FROM events WHERE date BETWEEN $1 AND $2 AND feature = $3"

	createIndex = `INSERT INTO index(event_id, features) VALUES ($1, $2) RETURNING id`
	getIndex    = `SELECT event_id, features FROM index WHERE id = $1`
	getFeatures = `SELECT features FROM idnex WHERE id = $1 `

	getCachedIds    = `SELECT * FROM cache`
	saveCache       = `INSERT INTO cache(id) VALUES($1)`
	deleteCache     = `DELETE FROM cache WHERE id = $1`
	isAlreadyCached = `SELECT * FROM cache WHERE id = $1`
)
