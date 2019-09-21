package randata

import (
	"context"
	"errors"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// RandomSSN will try to generate a valid random SSN by generating up until
// the given retries or until it generates a valid SSN whichever comes first.
// The retires default to 100 if no value is given.
// If formatted it will return a string with the format XXX-XX-XXXX opposed to
// non formatted XXXXXXXXX.
func RandomSSN(formatted bool, routines int) string {
	validSSN := make(chan string)

	if routines == 0 {
		routines = 100
	}

	min := 100000000
	max := 999999999

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	wg := sync.WaitGroup{}
	for i := 0; i < routines; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:
				r := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
				rand.Seed(time.Now().UTC().UnixNano())
				ssn := r.Intn(max-min+1) + min
				ssnStr := strconv.Itoa(ssn)
				valid, _ := validateSSN(ssnStr)
				if valid {
					cancel()
					validSSN <- ssnStr

				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(validSSN)
	}()

	var result string
	for ssn := range validSSN {
		result = ssn
		break
	}
	return result
}

func validateSSN(ssn string) (bool, error) {
	if len(ssn) != 9 {
		return false, errors.New("SSN length must be equal to 9")
	}

	if string(ssn[0]) == "0" {
		return false, errors.New("can not start with zero")
	}

	if ssn[0:2] == "000" || ssn[3:4] == "00" || ssn[5:8] == "0000" {
		return false, errors.New("no group can all be zeroes")
	}

	ssnInt, err := strconv.Atoi(ssn)
	if err != nil {
		return false, err
	}

	if ssnInt < 987654329 && ssnInt > 987654320 {
		return false, errors.New("SSN can not be in the range of 987654320-987654329")
	}
	return true, nil
}
