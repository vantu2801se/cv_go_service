module handlers

go 1.16

replace context/application/cvservice/api => ../api

require (
	context/application/cvservice/api v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.0
)

replace context/application/cvservice/models => ../models

replace context/application/cvservice/repositories => ../repositories

replace context/application/cvservice/entities => ../entities
