package randata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var (
	// latRegex is used to match a latitude.
	latRegex  = regexp.MustCompile(`^(\+|-)(?:90(?:(?:\.0{1,13})?)|(?:[0-9]|[1-8][0-9])(?:(?:\.[0-9]{1,13})?))$`)
	longRegex = regexp.MustCompile(`^(\+|-)?(?:180(?:(?:\.0{1,13})?)|(?:[0-9]|[1-9][0-9]|1[0-7][0-9])(?:(?:\.[0-9]{1,13})?))$`)

	jsonValuePrefix = []byte(`"`)
)

// Coordinate ...
type Coordinate float64

// ToFloat64 ...
func (b Coordinate) ToFloat64() float64 {
	return float64(b)
}

// ToString ...
func (b Coordinate) ToString() string {
	return strconv.FormatFloat(b.ToFloat64(), 'f', -1, 64)
}

// MarshalString ...
func (c *Coordinate) MarshalString(s string) error {
	i, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return err
	}

	*c = Coordinate(i)
	return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface, which allows us to
// ingest values of any json type as an int64 and run our custom conversion
func (c *Coordinate) UnmarshalJSON(data []byte) error {
	if !bytes.HasPrefix(data, jsonValuePrefix) {
		return json.Unmarshal(data, (*float64)(c))
	}
	var coordinateStr string
	if err := json.Unmarshal(data, &coordinateStr); err != nil {
		return err
	}

	if latRegex.MatchString(coordinateStr) || longRegex.MatchString(coordinateStr) {
		coordinates := strings.Split(coordinateStr, ".")
		if len(coordinates[1]) > 6 {
			coordinateStr = coordinates[1][:6]
		}
		return c.MarshalString(coordinateStr)
	}

	return fmt.Errorf("can not unmarshal string: %v, not a validate coordinate", coordinateStr)
}

type Coordinates struct {
	Latitude  Coordinate `json:"latitude,omitempty"`
	Longitude Coordinate `json:"longitude,omitempty"`
}

// Address represents an address that is loaded from an address file that uses
// the Starbucks locations in the US. See Google Autocomplete for more
// information on the naming of each field.
type Address struct {
	Coordinates
	Locality                 string `json:"locality"` // city
	Country                  string `json:"country"`
	StreetNumber             string `json:"street_number"`               // e.g., 496 ...
	UnitNumber               string `json:"unit_number"`                 // e.g., apt/unit...
	Route                    string `json:"route"`                       // street name
	PostalCode               string `json:"postal_code"`                 // zip code
	AdministrativeAreaLevel1 string `json:"administrative_area_level_1"` // state
}

// EmptyPostalCode ...
func (a Address) EmptyPostalCode() bool {
	return a.PostalCode == ""
}
