module example.com/albums

go 1.21.0

replace example.com/dbConnect => ../dbConnect

replace example.com/goDotEnvVariable => ../goDotEnvVariable

require (
	example.com/dbConnect v0.0.0-00010101000000-000000000000
	github.com/lib/pq v1.10.9
)

require (
	example.com/goDotEnvVariable v0.0.0-00010101000000-000000000000 // indirect
	github.com/joho/godotenv v1.5.1 // indirect
)
