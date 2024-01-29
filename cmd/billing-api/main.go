package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"

	"github.com/gin-gonic/gin"
	"github.com/jacobbeck/currency-converter/api"
	"github.com/jacobbeck/currency-converter/pkg/user"
)


type Router struct {
	engine *gin.Engine
}



func main() {
	db := connectToDB()
	initDB(db)

	defer db.Close()

	userService := user.NewService(db)
	
	r := &Router{}
	r.engine = gin.Default()

	r.engine.Use(api.UserMiddleware(*userService))

	// Define routes
	r.engine.GET("/users/:id/balance", api.GetUserBalanceHandler)
	r.engine.POST("/users", api.CreateUserHandler)

	// Run the server
	port := ":8080"
	fmt.Printf("Server running on %s\n", port)
	r.engine.Run(port)
}


func connectToDB() *sql.DB {
	host     := os.Getenv("PG_HOST")
	port     := 5432
	user     := os.Getenv("PG_USER")
	password := os.Getenv("PG_PASSWORD")
	dbname   := os.Getenv("PG_DB")

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// Open a database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	// Ping the database to check if the connection is successful
	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging the database:", err)
	}

	fmt.Println("Successfully connected to the PostgreSQL database!")

	return db
}

func initDB(db *sql.DB) {
	const (
		createUserTable = `
		CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
			name TEXT,
			balance INTEGER
		)
		`
	)

	_, err := db.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\";")
	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec(createUserTable)
	if err != nil {
		panic("failed to initialize User table")
	}
}