package validation

import (
	"database/sql"
	"errors"
	"net/http"
	"renthouses/schema"

	DBconfig "github.com/sridharu346/newrepo/dbconfig"
)

func addHousevalidation(db *sql.DB, newHouse schema.HouseDetails) (int, error) {
	if !IsValidAddress(newHouse.Address) {
		return http.StatusBadRequest, errors.New("Address is not valid")
	}
	if DBconfig.IsHouseExist(db, newHouse.Address) {
		return http.StatusForbidden, errors.New("House already exists")
	}

	if IsValidBedRooms(newHouse.BedRooms) != nil {
		return http.StatusBadRequest, errors.New("Invalid number of bedRooms")
	}

	if IsValidBathrooms(newHouse.BathRooms) != nil {
		return http.StatusBadRequest, errors.New("Invalid number of bathrooms")
	}

	if IsRentValid(newHouse.RentCost) != nil {
		return http.StatusBadRequest, errors.New("Invalid rentCost")
	}

	if IsLeasingPeriod(newHouse.LeasingPeriod) != nil {
		return http.StatusBadRequest, errors.New("Lease period should not be empty")
	}
}
