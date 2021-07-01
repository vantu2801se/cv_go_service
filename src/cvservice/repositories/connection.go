package repositories

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
)

// type conn struct {
// 	dbType   string
// 	username string
// 	password string
// 	uri      string
// 	port     int
// 	database string
// }

func getConnection(dbType string, username string, password string, uri string, port int, database string) (*sql.DB, error) {
	//"{USERNAME}:{PASSWORD}@tcp({URI}:{PORT})/{DATABASE}"
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, uri, port, database)
	return sql.Open(dbType, url)
}

func GetMySqlConnection(username string, password string, uri string, port int, database string) (*sql.DB, error) {
	return getConnection("mysql", username, password, uri, port, database)
}
func GetPosgresqlConnection(username string, password string, uri string, port int, database string) (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=require",
		uri, port, username, password, database)
	return sql.Open("postgres", psqlInfo)
}
