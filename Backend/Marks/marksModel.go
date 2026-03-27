package marks

import (
	"errors"
	"fmt"

	"github.com/Anuj700077/Dummy-project/database"
)

type Marks struct {
	ID         int64   `json:"id"`
	Sid        int64   `json:"sid"`
	Sname      string  `json:"sname"`
	Math       int64   `json:"math"`
	Science    int64   `json:"science"`
	Hindi      int64   `json:"hindi"`
	English    int64   `json:"english"`
	Computer   int64   `json:"computer"`
	Total      int64   `json:"total"`
	Percentage float64 `json:"percentage"`
}

// CREATE OR UPSERT
func (m *Marks) CreateMark() error {

	if m.Sid == 0 {
		return errors.New("student id is required")
	}

	m.Total = m.Math + m.Science + m.Hindi + m.English + m.Computer
	m.Percentage = float64(m.Total) / 5

	_, err := database.DB.Exec(
		`INSERT INTO marks 
		(student_id, math, science, hindi, english, computer, total, percentage) 
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (student_id) 
		DO UPDATE SET 
			math=$2, science=$3, hindi=$4, english=$5, computer=$6, total=$7, percentage=$8`,
		m.Sid, m.Math, m.Science, m.Hindi, m.English, m.Computer, m.Total, m.Percentage,
	)

	if err != nil {
		fmt.Println("DB Insert Error:", err)
		return errors.New("data not inserted")
	}

	return nil
}

// GET ALL
func GetAllMarks() ([]Marks, error) {

	query := `
	SELECT m.id, m.student_id, s.sname, m.math, m.science, m.hindi, m.english, m.computer, m.total, m.percentage
	FROM marks m
	INNER JOIN students s ON m.student_id = s.id
	ORDER BY m.id;
	`

	rows, err := database.DB.Query(query)
	if err != nil {
		fmt.Println("Query Error:", err)
		return nil, errors.New("failed to fetch data")
	}
	defer rows.Close()

	var marksList []Marks

	for rows.Next() {
		var m Marks
		err := rows.Scan(
			&m.ID, &m.Sid, &m.Sname, &m.Math, &m.Science, &m.Hindi, &m.English, &m.Computer, &m.Total, &m.Percentage,
		)
		if err != nil {
			fmt.Println("Scan Error:", err)
			return nil, errors.New("error reading data")
		}
		marksList = append(marksList, m)
	}

	if err = rows.Err(); err != nil {
		return nil, errors.New("error processing data")
	}

	return marksList, nil
}

// UPDATE
func (m *Marks) UpdateMark() error {

	if m.Sid == 0 {
		return errors.New("student id required")
	}

	m.Total = m.Math + m.Science + m.Hindi + m.English + m.Computer
	m.Percentage = float64(m.Total) / 5

	result, err := database.DB.Exec(
		`UPDATE marks SET 
			math=$1,
			science=$2,
			hindi=$3,
			english=$4,
			computer=$5,
			total=$6,
			percentage=$7
		WHERE student_id=$8`,
		m.Math, m.Science, m.Hindi, m.English, m.Computer, m.Total, m.Percentage, m.Sid,
	)
	if err != nil {
		return errors.New("data not updated")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no record found to update")
	}

	return nil
}

// DELETE
func DeleteMark(sid int64) error {

	result, err := database.DB.Exec(
		"DELETE FROM marks WHERE student_id=$1",
		sid,
	)

	if err != nil {
		return errors.New("data not deleted")
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return errors.New("no record found to delete")
	}

	return nil
}
