module repositories

go 1.16

replace context/application/cvservice/entities => ../entities

require (
	context/application/cvservice/entities v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.6.0
	github.com/lib/pq v1.10.2
)
