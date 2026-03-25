package models

import (
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

const MonthlyFee = 5000

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

	// check previous due
	var lastDue int64
	err = database.DB.QueryRow(`
		SELECT amtdue FROM fees 
		WHERE student_id=$1 
		ORDER BY feemonth DESC 
		LIMIT 1
	`, f.Sid).Scan(&lastDue)

	if err != nil && err.Error() != "sql: no rows in result set" {
		return errors.New("error checking previous dues")
	}

	if lastDue > 0 {
		return fmt.Errorf("previous due pending: %d", lastDue)
	}

	calculatedDue := MonthlyFee - f.Amtpaid
	if calculatedDue < 0 {
		calculatedDue = 0
	}

	query := `
	INSERT INTO fees (student_id, feemonth, amtpaid, amtdue, receivedate)
	VALUES ($1, $2, $3, $4, $5)
	`
	_, err = database.DB.Exec(
		query,
		f.Sid,
		f.FeeMonth,
		f.Amtpaid,
		calculatedDue,
		f.ReceiveDate,
	)

	if err != nil {
		fmt.Println("DB ERROR:", err)
		return errors.New("fees not inserted")
	}

	return nil
}

func GetAllFees() ([]Fees, error) {

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
		ORDER BY f.feemonth DESC
	`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var fees []Fees

	for rows.Next() {
		var f Fees
		err := rows.Scan(
			&f.ID, &f.Sid, &f.Sname, &f.Fname, &f.FeeMonth, &f.Amtpaid, &f.Amtdue, &f.ReceiveDate,
		)
		if err != nil {
			return nil, err
		}
		fees = append(fees, f)
	}

	return fees, nil
}

func UpdateFess(){
	
}
