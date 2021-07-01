module context/application

go 1.16

replace context/application/cvservice/handlers => /cvservice/handlers

replace context/application/cvservice/api => /cvservice/api

replace context/application/cvservice/models => /cvservice/models

replace context/application/cvservice/repositories => /cvservice/repositories

replace context/application/cvservice/entities => /cvservice/entities

require context/application/cvservice/handlers v0.0.0-00010101000000-000000000000
