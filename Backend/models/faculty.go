package models

import "github.com/Anuj700077/Dummy-project/database"

type Faculty struct {
	ID         int64  `json:"id"`
	Tname      string `json:"tname"`
	Subject    string `json:"subject"`
	Department string `json:"department"`
	DOA        string `json:"doa"`
}

func (f *Faculty) CreateFaculty() error {
	_, err := database.DB.Exec(
		"INSERT INTO faculty (tname, subject, department, doa) VALUES($1, $2, $3, $4)",
		f.Tname, f.Subject, f.Department, f.DOA,
	)
	return err

}

func GetAllFaculty() ([]Faculty, error) {
	rows, err := database.DB.Query("SELECT id, tname, subject, department, doa FROM faculty")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var faculty []Faculty

	for rows.Next() {
		var f Faculty
		err := rows.Scan(&f.ID, &f.Tname, &f.Subject, &f.Department, &f.DOA)
		if err != nil {
			return nil, err
		}
		faculty = append(faculty, f)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return faculty, nil
}

func (f *Faculty) UpdateFaculty(id int64) error {
	_, err := database.DB.Exec(
		"UPDATE faculty SET tname=$1, subject=$2, department=$3, doa=$4 WHERE id=$5",
		f.Tname, f.Subject, f.Department, f.DOA, id,
	)
	return err
}

func DeleteFaculty(id int64) error {
	_, err := database.DB.Exec("DELETE FROM faculty WHERE id=$1", id)
	return err
}

