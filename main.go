package main

import (
	"fmt"
	"net/http"

	DBconfig "github.com/sridharu346/newrepo/dbconfig"
)

func main() {
	DB, err := DBconfig.DBConnection()
	if err != nil {
		fmt.Print("Error Occured")
	}
	defer DB.Close()
	fmt.Println("DB connected Successfully")

	//hosting the server
	fmt.Println("Local host is servered at port 8800")
	err = http.ListenAndServe(":8800", nil)
	if err != nil {
		fmt.Println("Error in hosting the server", err)
	}


}