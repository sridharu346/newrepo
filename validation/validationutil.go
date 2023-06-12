package validation

import (
	"errors"
	"regexp"
)

func IsValidAddress(address string) bool {
	// Define the regular expression pattern for the address format
	pattern := `^\d+\s+[\w\s]+\s*,\s*\w+\s*,\s*\w+\s+\d{5}$`
	//example: 123 Main St, City, State 12345

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Check if the address matches the pattern
	return regex.MatchString(address)
}

func IsValidBedRooms(numBedrooms int) error {
	if numBedrooms < 0 {
		return errors.New("Invalid number of bedrooms")
	}
	return nil
}

func IsValidBathrooms(numBathrooms int) error {
	if numBathrooms < 0 {
		return errors.New("Invalid number of bathrooms")

	}
	return nil
}

func IsRentValid(rentCost int) error {
	if rentCost < 0 {
		return errors.New("Invalid rentCost")
	}
	return nil
}

func IsLeasingPeriod(leasingperiod string) error {
	if leasingperiod == "" {
		return errors.New("Lease period should not be empty")
	}
	return nil
}
