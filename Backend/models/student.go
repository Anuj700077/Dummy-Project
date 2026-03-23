package models

import "github.com/Anuj700077/Dummy-project/database"

type Student struct {
	ID      int64  `json:"id"`
	Sname   string `json:"sname"`
	Fname   string `json:"fname"`
	Address string `json:"address"`
	Dob     string `json:"dob"`
}

// Here Create Student
func (s *Student) Create() error {
	_, err := database.DB.Exec(
		"INSERT INTO students(sname, fname, address, dob) VALUES($1,$2,$3,$4)",
		s.Sname,
		s.Fname,
		s.Address,
		s.Dob,
	)
	return err
}

// Here fetch  All Students
func GetAllStudents() ([]Student, error) {
	rows, err := database.DB.Query("SELECT id, sname, fname, address, dob FROM students")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var s Student
		rows.Scan(&s.ID, &s.Sname, &s.Fname, &s.Address, &s.Dob)
		students = append(students, s)
	}

	return students, nil
}

// Here update Student
func (s *Student) Update(id string) error {
	_, err := database.DB.Exec(
		`UPDATE students SET sname=$1, fname=$2, address=$3, dob=$4 WHERE id=$5`,
		s.Sname, s.Fname, s.Address, s.Dob, id,
	)
	return err
}

// here Delete Student
func DeleteStudentByID(id string) error {
	_, err := database.DB.Exec("DELETE FROM students WHERE id=$1", id)
	return err
}
