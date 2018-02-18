package testdata

import (
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

// Address represents an address that is loaded from an address file that uses
// all the starbucks location in the US. See gogole autplace complete for more
// information on the naming of each fields
type Address struct {
	Locality                 string `json:"locality"` // city
	Country                  string `json:"country"`
	Latitude                 string `json:"latitude"`
	Longitude                string `json:"longitude"`
	StreetNumber             string `json:"street_number"`               // 496 ...
	UnitNumber               string `json:"unit_number"`                 // apt/unit...
	Route                    string `json:"route"`                       // street name
	PostalCode               string `json:"postal_code"`                 // zip code
	AdministrativeAreaLevel1 string `json:"administrative_area_level_1"` // state
}

// USAddresses ...
var USAddresses = make([]Address, 0)

func init() {
	data, err := Asset("data/us_addresses.json")
	if err != nil {
		log.Fatalf("fatal error loading testdata, error: %s", err.Error())
	}
	err = json.Unmarshal(data, &USAddresses)
	if err != nil {
		panic(err)
	}
}

// RandomUSAddress ...
func RandomUSAddress() Address {
	src := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(src)
	return USAddresses[rnd.Intn(len(USAddresses))]
}
