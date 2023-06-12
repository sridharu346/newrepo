package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"renthouses/schema"
)

func addHouse(w http.ResponseWriter, r *http.Request) {

	//give the details
	if r.Method != "POST" {
		ResponseFormat(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//DB CONNECTION
	db, err := DBconfig.DBConnection()
	if err != nil {
		ResponseFormat(w, "DB connection failed", http.StatusInternalServerError)
	}
	defer db.Close()

	//reading the body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// unmarshal the body
	var newDetails schema.HouseDetails
	err = json.Unmarshal(body, &newDetails)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//validation

	validate, err := validation.addHousevalidation(db, newHouse)
	if err != nil {
		ResponseFormat(w, err.Error(), validate)

	}

	//pushing data
	ErrorCode, err := DBconfig.InsertHouse(db, newHouse)
	if err != nil {
		ResponseFormat(w, err.Error(), ErrorCode)
		return
	}
	ResponseFormat(w, fmt.Sprintf("House %s created successfully", newHouse.Address), http.StatusCreated, nil)
}
