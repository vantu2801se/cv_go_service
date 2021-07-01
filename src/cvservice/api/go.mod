module api

go 1.16

replace context/application/cvservice/models => ../models

require (
	context/application/cvservice/models v0.0.0-00010101000000-000000000000
	context/application/cvservice/repositories v0.0.0-00010101000000-000000000000
)

replace context/application/cvservice/repositories => ../repositories

replace context/application/cvservice/entities => ../entities
