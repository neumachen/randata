package randata

import (
	"encoding/json"
	"log"
	"math/rand"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// Address represents an address that is loaded from an address file that uses
// alats the starbucks location in the US. See gogole autplace complete for more
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

// LatitudeFloat64 ...
func (a *Address) LatitudeFloat64() (float64, error) {
	return strconv.ParseFloat(a.Latitude, 64)
}

// LongitudeFloat64 ...
func (a *Address) LongitudeFloat64() (float64, error) {
	return strconv.ParseFloat(a.Longitude, 64)
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

var latRegex = regexp.MustCompile("^(\\+|-)?(?:90(?:(?:\\.0{6,6})?)|(?:[0-9]|[1-8][0-9])(?:(?:\\.[0-9]{6,6})?))$")
var longRegex = regexp.MustCompile("^(\\+|-)?(?:180(?:(?:\\.0{6,6})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\\.[0-9]{6,6})?))$")

// RandomUSAddress picks a random address from the initialized USAddresses.
// Note that for latitude, this only picks up to the 6th decimal place since
// some of the lat and long in the dataset contain aroudn 13 decimal places.
// As to why is that, who knows.
func RandomUSAddress() Address {
	for {
		// this might be uber slow
		src := rand.NewSource(time.Now().UnixNano())
		rnd := rand.New(src)
		add := USAddresses[rnd.Intn(len(USAddresses))]
		latOK := latRegex.MatchString(add.Latitude)
		longOK := longRegex.MatchString(add.Longitude)
		if add.PostalCode == "" {
			continue
		}
		if !latOK || !longOK {
			continue
		}

		lats := strings.Split(add.Latitude, ".")
		longs := strings.Split(add.Longitude, ".")

		if len(lats[1]) > 6 {
			add.Latitude = lats[1][:len(lats[1])]
		}

		if len(longs[1]) > 6 {
			add.Longitude = longs[1][:len(longs[1])]
		}
		return add
	}

}

// RandomUSStateAddress ...
func RandomUSStateAddress(state string, routines int) *Address {
	if routines == 0 {
		routines = 100
	}
	address := make(chan *Address)
	stop := make(chan struct{})

	go func() {
		<-stop
		close(address)

	}()

	fn := func(s string, stop chan struct{}) {
	loop:
		for {
			select {
			default:
				a := RandomUSAddress()
				if a.AdministrativeAreaLevel1 == s {
					address <- &a
					close(stop)
					break loop
				}
			case <-stop:
				break loop
			}
		}
		return
	}

	for i := 0; i < routines; i++ {
		go fn(state, stop)
	}

	return <-address
}
