package database

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB() {
	godotenv.Load()

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

	collegeDB := os.Getenv("DB_NAME")

	_, err = db.Exec("CREATE DATABASE " + collegeDB)
	if err != nil {
		fmt.Println("Database already exists or no permission")
	} else {
		fmt.Println("✅ Database created")
	}

	db.Close()

	actualConn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		collegeDB,
	)

	DB, err = sql.Open("postgres", actualConn)
	if err != nil {
		panic(err)
	}

	err = DB.Ping()
	if err != nil {
		panic("❌ DB not reachable")
	}

	fmt.Println("✅ Connected to college_db")
}

func CreateTable() {
	query := `
	CREATE TABLE IF NOT EXISTS students (
		id SERIAL PRIMARY KEY,
		sname TEXT NOT NULL,
		fname TEXT NOT NULL,
		address TEXT NOT NULL,
		dob DATE NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Students Table Created")
}

func CreateFacultyTable() {
	query := `
	CREATE TABLE IF NOT EXISTS faculty (
		id SERIAL PRIMARY KEY,
		tname TEXT NOT NULL,
		subject TEXT NOT NULL,
		department TEXT NOT NULL,
		doa DATE NOT NULL
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Faculty Table Created")
}

func CreateMarksTable() {
	query := `
	CREATE TABLE IF NOT EXISTS marks (
		id SERIAL PRIMARY KEY,
		student_id INT UNIQUE NOT NULL,

		math INT DEFAULT 0,
		science INT DEFAULT 0,
		hindi INT DEFAULT 0,
		english INT DEFAULT 0,
		computer INT DEFAULT 0,

		total INT DEFAULT 0,
		percentage FLOAT DEFAULT 0,

		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);
	`

	_, err := DB.Exec(query)
	if err != nil {
		panic(err)
	}

	fmt.Println("✅ Marks Table Created")
}

func CreateFeeTable() {
	query := `
	CREATE TABLE IF NOT EXISTS fees (
		id SERIAL PRIMARY KEY,
		student_id INT NOT NULL,
		feemonth DATE NOT NULL,
		amtpaid INT NOT NULL,
		amtdue INT NOT NULL,
		receivedate DATE NOT NULL,

		
		UNIQUE(student_id, feemonth),

		FOREIGN KEY (student_id) REFERENCES students(id) ON DELETE CASCADE
	);

	CREATE INDEX IF NOT EXISTS idx_fees_student_id ON fees(student_id);
	`

	_, err := DB.Exec(query)
	if err != nil {
		fmt.Println(" Could not create Fees table:", err)
		return
	}

	fmt.Println("✅ Fees table created successfully")
}
