package lib

import "github.com/taubyte/go-sdk/database"

func open() (database.Database, error) {
	return database.New(databaseMatch)
}
