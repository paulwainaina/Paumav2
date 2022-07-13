module example.com/pauma

go 1.18

replace example.com/controllers => ./controllers

replace example.com/models => ./models

require (
	example.com/controllers v0.0.0-00010101000000-000000000000
	github.com/joho/godotenv v1.4.0
)

require example.com/models v0.0.0-00010101000000-000000000000 // indirect
