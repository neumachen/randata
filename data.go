package randata

import (
	"embed"
	"encoding/json"
)

//go:embed data/us_addresses.json
var data embed.FS

// USAddresses stores the addresses loaded from the JSON data.
var USAddresses = make([]Address, 0)

func init() {
	b, err := data.ReadFile("data/us_addresses.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(b, &USAddresses)
	if err != nil {
		panic(err)
	}
}
