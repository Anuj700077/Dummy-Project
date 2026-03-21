package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	godotenv.Load()

	// Step 1: Connect to default postgres DB
	defaultConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
	)

	db, err := sql.Open("postgres", defaultConn)
	if err != nil {
		panic(err)
	}

	// Step 2: Create database if not exists
	college_db := os.Getenv("DB_NAME")
	_, err = db.Exec("CREATE DATABASE " + college_db)
	if err != nil {
		fmt.Println(" Database may already exist")
	} else {
		fmt.Println("✅ Database created")
	}

	db.Close()

	// Step 3: Connect to your actual DB
	actualConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		college_db,
	)

	DB, err = sql.Open("postgres", actualConn)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Connected to college_db")
}


func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		sname TEXT,
		fname TEXT,
		address TEXT,
		dob DATE
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Table Created")
}
