Site on GoLang. 

Save and manage your todo's marks. 

To start:

1. Run docker image with postgres 
-- sudo docker run --name some-postgres -e POSTGRES_PASSWORD=postgres -d -p 5432:5432 postgres
2. Exec sql query from db.sql to create todo_db.
3. Check config.json 
4. Run server from soure by command: go run main.go config config.json