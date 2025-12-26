DB_URL=postgres://username:password@localhost:5432/ostadbun?sslmode=disable

create:
	@migrate create -ext sql -dir migrations -seq $(dis)
up:
	migrate -path database/migrations -database "$(DB_URL)" up

down:
	migrate -path database/migrations -database "$(DB_URL)" down 1

run:
	DATABASE_URL="$(DB_URL)"  go run main.go