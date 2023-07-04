package randata

import (
	"context"
	"crypto/rand"
	"math/big"

	"golang.org/x/sync/errgroup"
)

// USAddress picks a random address from the initialized USAddresses.
// Note that for latitude, this only picks up to the 6th decimal place since
// some of the lat and long in the dataset contain around 13 decimal places.
// The reason for this limitation is unknown.
func USAddress() (*Address, error) {
	count := big.NewInt(int64(len(USAddresses)))
	for {
		randIndex, err := rand.Int(rand.Reader, count)
		if err != nil {
			return nil, err
		}
		randomAddress := USAddresses[int(randIndex.Int64())]

		if randomAddress.EmptyPostalCode() {
			continue
		}

		return &randomAddress, nil
	}
}

// USStateAddress returns a random address from the initialized USAddresses
// that belongs to the specified state. It uses multiple goroutines to improve
// performance, with the number of goroutines specified by the 'routines' parameter.
// If 'routines' is 0, it defaults to 10.
func USStateAddress(ctx context.Context, state string, routines int) (*Address, error) {
	if routines == 0 {
		routines = 15
	}

	localCtx, done := context.WithCancel(ctx)
	defer done()

	errGroup, gCtx := errgroup.WithContext(localCtx)
	addresses := make(chan *Address)
	for i := 0; i < routines; i++ {
		errGroup.Go(func() error {
		loop:
			for {
				select {
				case <-gCtx.Done():
					return gCtx.Err()
				default:
					a, err := USAddress()
					if err != nil {
						return err
					}
					if a.AdministrativeAreaLevel1 == state {
						addresses <- a
						done()
						break loop
					}
				}
			}

			return gCtx.Err()
		})
	}

	errChan := make(chan error, 1)
	go func() {
		defer func() {
			close(addresses)
			close(errChan)
		}()
		if err := errGroup.Wait(); err == nil || err == context.Canceled {
			return
		} else {
			errChan <- err
		}
	}()

	return <-addresses, <-errChan
}
