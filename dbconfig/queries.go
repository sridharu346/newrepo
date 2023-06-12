package dbconfig

import (
	"database/sql"
	"net/http"
	"renthouses/schema"
)

func IsHouseExist(db *sql.DB, address string) bool {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM rentalhouses WHERE Address = ?", address).Scan(&count)
	if err != nil {
		return false
	}
	if count > 0 {
		return true
	}
	return false
}

func InsertHouse(db *sql.DB, newHouse schema.HouseDetails) (int error) {
	stmt, err := db.Prepare(`Insert into rentalhouses(Bedrooms,Bathrooms,Address,rentcost,leasingperiod)values(?,?,?,?,?)`)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	_, err = stmt.Exec(newHouse.Bedrooms, newHouse.Bathrooms, newHouse.Address, newHouse.rentcost, newHouse.leasingperiod)
	if err != nil {
		return http.StatusInternalServerError, err
	}
	return http.Statuscreated, nil
}
