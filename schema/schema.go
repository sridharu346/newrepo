package schema

type HouseDetails struct {
	BedRooms      int    `json:"BedRooms"`
	BathRooms     int    `json:"BathRooms"`
	Address       string `json:"Address"`
	RentCost      int    `json:"RentCost"`
	LeasingPeriod string `json:"LeasingPeriod"`
}
