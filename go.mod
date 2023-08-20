module example/data-access

go 1.21.0

require (
	example.com/goDotEnvVariable v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require github.com/joho/godotenv v1.5.1 // indirect

replace example.com/goDotEnvVariable => ./goDotEnvVariable
