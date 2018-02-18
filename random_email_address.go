package randata

import (
	"strings"

	"github.com/magicalbanana/tg"
)

// RandomEmailAddress ...
func RandomEmailAddress(domain string, userLength int) string {
	if userLength == 0 {
		userLength = 10
	}
	user, _ := tg.RandGen(userLength, tg.LowerUpperDigit, "", "")
	return strings.Join([]string{user, domain}, "@")
}
