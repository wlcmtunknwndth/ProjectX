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
                   			feature,
							city,
                   			address,
							name,
							img_path,
							description
                   			) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)	
	`
	getId         = "SELECT id FROM events WHERE name = $1 AND date = $2"
	changeImgPath = "UPDATE events SET img_path=$1"
	patchEvent    = `UPDATE events SET price = $1,
											restrictions = $2,
											date = $3,
											feature = $4,
											city = $5,
											address = $6,
											name = $7,
											description = $8
									WHERE id = $9
											`

	deleteEvent        = "DELETE FROM events WHERE id = $1"
	getEventsByFeature = "SELECT * FROM events WHERE date BETWEEN $1 AND $2 AND feature = $3"

	getCachedIds    = `SELECT * FROM cache`
	saveCache       = `INSERT INTO cache(id) VALUES($1)`
	deleteCache     = `DELETE FROM cache WHERE id = $1`
	isAlreadyCached = `SELECT * FROM cache WHERE id = $1`
)
