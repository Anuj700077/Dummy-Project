package models

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/Anuj700077/Dummy-project/database"
)

type Fees struct {
	ID          int64  `json:"id"`
	Sid         int64  `json:"sid"`
	Sname       string `json:"sname"`
	Fname       string `json:"fname"`
	FeeMonth    string `json:"feemonth"`
	Amtpaid     int64  `json:"amtpaid"`
	Amtdue      int64  `json:"amtdue"`
	ReceiveDate string `json:"receivedate"`
}

const MonthlyFee int64 = 5000


func CreateFee(f Fees) error {


	var exists int
	err := database.DB.QueryRow(
		"SELECT COUNT(*) FROM students WHERE id=$1", f.Sid,
	).Scan(&exists)

	if err != nil {
		return errors.New("could not fetch student")
	}
	if exists == 0 {
		return errors.New("student not found")
	}

	
	var lastDue int64
	var lastMonth string

	err = database.DB.QueryRow(`
		SELECT feemonth, amtdue FROM fees 
		WHERE student_id=$1 
		ORDER BY id DESC 
		LIMIT 1
	`, f.Sid).Scan(&lastMonth, &lastDue)

	if err != nil {
		if err == sql.ErrNoRows {
			lastDue = 0
			lastMonth = ""
		} else {
			return errors.New("error fetching last record")
		}
	}

	
	if lastMonth == f.FeeMonth {

		totalFee := lastDue
		calculatedDue := totalFee - f.Amtpaid

		if calculatedDue < 0 {
			calculatedDue = 0
		}

		_, err = database.DB.Exec(`
			UPDATE fees 
			SET amtpaid = amtpaid + $1,
			    amtdue = $2,
			    receivedate = $3
			WHERE student_id=$4 AND feemonth=$5
		`,
			f.Amtpaid,
			calculatedDue,
			f.ReceiveDate,
			f.Sid,
			f.FeeMonth,
		)

		if err != nil {
			return errors.New("failed to update fee")
		}

		return nil
	}


	if lastDue > 0 {
		return fmt.Errorf(" clear previous due first: %d", lastDue)
	}


	totalFee := MonthlyFee               
	calculatedDue := totalFee - f.Amtpaid 

	if calculatedDue < 0 {
		calculatedDue = 0
	}

	_, err = database.DB.Exec(`
		INSERT INTO fees (student_id, feemonth, amtpaid, amtdue, receivedate)
		VALUES ($1, $2, $3, $4, $5)
	`,
		f.Sid,
		f.FeeMonth,
		f.Amtpaid,
		calculatedDue,
		f.ReceiveDate,
	)

	if err != nil {
		return errors.New("fees not inserted")
	}

	return nil
}

func GetLatestFees() ([]Fees, error) {

	rows, err := database.DB.Query(`
		SELECT DISTINCT ON (f.student_id)
			f.id,
			f.student_id,
			s.sname,
			s.fname,
			f.feemonth,
			f.amtpaid,
			f.amtdue,
			f.receivedate
		FROM fees f
		JOIN students s ON s.id = f.student_id
		ORDER BY f.student_id, f.id DESC
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fees []Fees

	for rows.Next() {
		var f Fees
		err := rows.Scan(
			&f.ID, &f.Sid, &f.Sname, &f.Fname,
			&f.FeeMonth, &f.Amtpaid, &f.Amtdue, &f.ReceiveDate,
		)
		if err != nil {
			return nil, err
		}
		fees = append(fees, f)
	}

	return fees, nil
}


func GetFeesByStudentID(sid int64) ([]Fees, error) {

	rows, err := database.DB.Query(`
		SELECT 
			f.id,
			f.student_id,
			s.sname,
			s.fname,
			f.feemonth,
			f.amtpaid,
			f.amtdue,
			f.receivedate
		FROM fees f
		JOIN students s ON s.id = f.student_id
		WHERE f.student_id = $1
		ORDER BY f.id DESC
	`, sid)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fees []Fees

	for rows.Next() {
		var f Fees
		err := rows.Scan(
			&f.ID, &f.Sid, &f.Sname, &f.Fname,
			&f.FeeMonth, &f.Amtpaid, &f.Amtdue, &f.ReceiveDate,
		)
		if err != nil {
			return nil, err
		}
		fees = append(fees, f)
	}

	return fees, nil
}
