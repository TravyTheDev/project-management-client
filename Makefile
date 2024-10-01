up: 
	@goose -dir=migrations sqlite3 app.db up

down: 
	@goose -dir=migrations sqlite3 app.db down