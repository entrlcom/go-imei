package imei_cd_model

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

// CD — Check Digit.
type CD string

func (x CD) Validate() error {
	if len(x) != 1 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return errors.New("invalid CD")
	}

	return nil
}

func NewCD(s string) (CD, error) {
	x := CD(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}

func ComputeCD(s string) (CD, error) {
	var sum int64 = 0

	for i, v := range strings.Split(s, "") {
		n, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return "", fmt.Errorf("strconv.ParseInt: %w", err)
		}

		if i%2 == 1 {
			n *= 2
		}

		if n > 9 { //nolint:mnd // OK.
			n = n%10 + 1 //nolint:mnd // OK.
		}

		sum += n
	}

	sum = (10 - sum%10) % 10 //nolint:mnd // OK.

	return NewCD(strconv.FormatInt(sum, 10))
}
