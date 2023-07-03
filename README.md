![main](https://github.com/neumachen/randata/actions/workflows/ci.yml/badge.svg?branch=main)
![ci](https://github.com/neumachen/randata/actions/workflows/ci.yml/badge.svg)
[![GoDoc](https://img.shields.io/badge/godoc-reference-blue.svg)](https://pkg.go.dev/github.com/neumachen/randata)

# randata

The `randata` package provides functionalities for generating random address and Social Security Number (SSN) data. It includes functions for generating random addresses and validating SSNs.

## Installation

To use the `randata` package in your Go project, you can install it using the `go get` command:

```shell
go get github.com/neumachen/randata
```

## Usage

Import the `randata` package in your Go code:

```go
import "github.com/neumachen/randata"
```

### Address Struct

The `Address` struct represents an address with various fields such as locality (city), country, latitude, longitude, street number, unit number, route (street name), postal code (zip code), and administrative area level 1 (state).

```go
type Address struct {
	Locality                 string `json:"locality"`
	Country                  string `json:"country"`
	LatitudeStr              string `json:"latitude"`
	LongitudeStr             string `json:"longitude"`
	StreetNumber             string `json:"street_number"`
	UnitNumber               string `json:"unit_number"`
	Route                    string `json:"route"`
	PostalCode               string `json:"postal_code"`
	AdministrativeAreaLevel1 string `json:"administrative_area_level_1"`
}
```

### Address Methods

The `Address` struct provides the following methods:

- `EmptyPostalCode() bool`: Checks if the postal code field is empty.
- `ValidLatitude() bool`: Validates the latitude field using a regular expression.
- `InvalidLatitude() bool`: Checks if the latitude field is invalid.
- `ValidLongitude() bool`: Validates the longitude field using a regular expression.
- `InvalidLongitude() bool`: Checks if the longitude field is invalid.
- `LatitudeFloat64() (float64, error)`: Converts the latitude string to a float64 value.
- `LongitudeFloat64() (float64, error)`: Converts the longitude string to a float64 value.

### RandomUSAddress

The `RandomUSAddress` function picks a random address from a pre-initialized list of US addresses. It ensures that the selected address has a non-empty postal code, valid latitude, and valid longitude. Note that the latitude and longitude values are rounded to the 6th decimal place.

```go
func RandomUSAddress() (*Address, error)
```

### RandomUSStateAddress

The `RandomUSStateAddress` function returns a random address from the initialized list of US addresses that belongs to the specified state. It utilizes multiple goroutines to improve performance, with the number of goroutines specified by the `routines` parameter. If `routines` is set to 0, it defaults to 15.

```go
func RandomUSStateAddress(ctx context.Context, state string, routines int) (*Address, error)
```

The `ctx` parameter allows you to pass a context to control the cancellation or timeout of the function.

### RandomSSN

The `RandomSSN` function generates a random Social Security Number (SSN). It tries to generate a valid SSN by generating random numbers up to the given `retries` or until it generates a valid SSN, whichever comes first. The `formatted` parameter determines whether the generated SSN should be formatted (XXX-XX-XXXX) or not (XXXXXXXXX). It defaults to non-formatted if no value is given.

```go
func RandomSSN(formatted bool, routines int) string
```

### ValidateSSN

The `ValidateSSN` function validates an SSN. It checks the length, leading zeros, all-zero groups, and whether it matches the SSN regular expression.

```go
func ValidateSSN(ss

n string) (bool, error)
```

## Examples

Here's an example that demonstrates how to use the `randata` package to generate random US addresses and SSNs:

```go
package main

import (
	"context"
	"fmt"

	"github.com/neumachen/randata"
)

func main() {
	// Generate a random US address
	address, err := randata.RandomUSAddress()
	if err != nil {
		fmt.Println("Failed to generate random address:", err)
		return
	}

	fmt.Println("Random Address:")
	fmt.Println("Locality:", address.Locality)
	fmt.Println("Country:", address.Country)
	fmt.Println("Latitude:", address.LatitudeStr)
	fmt.Println("Longitude:", address.LongitudeStr)
	fmt.Println("Street Number:", address.StreetNumber)
	fmt.Println("Unit Number:", address.UnitNumber)
	fmt.Println("Route:", address.Route)
	fmt.Println("Postal Code:", address.PostalCode)
	fmt.Println("Administrative Area Level 1:", address.AdministrativeAreaLevel1)

	// Generate a random SSN
	ssn := randata.RandomSSN(true, 100)
	fmt.Println("Random SSN:", ssn)
}
```

This example demonstrates how to generate a random US address and a random SSN using the `randata` package. The generated address and SSN are then printed to the console.
