package imei_cd

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidCD = errors.New("invalid CD")

// CD — Check Digit.
type CD string

func (x CD) IsEqual(v CD) bool {
	return x == v
}

func (x CD) IsValid() bool {
	return x.Validate() == nil
}

func (x CD) IsZero() bool {
	return len(x) == 0
}

func (x CD) String() string {
	return string(x)
}

func (x CD) Validate() error {
	if len(x) != 1 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return ErrInvalidCD
	}

	return nil
}

func NewCD(s string) (CD, error) {
	x := CD(s)

	if err := x.Validate(); err != nil {
		return "", err
	}

	return x, nil
}

func ComputeCD(s string) (CD, error) {
	var sum int64 = 0

	for i, v := range strings.Split(s, "") {
		n, err := strconv.ParseInt(v, 10, 8)
		if err != nil {
			return "", errors.Join(err, ErrInvalidCD)
		}

		if i%2 == 1 {
			n *= 2
		}

		if n > 9 { //nolint:gomnd // OK.
			n = n%10 + 1 //nolint:gomnd // OK.
		}

		sum += n
	}

	sum = (10 - sum%10) % 10 //nolint:gomnd // OK.

	return NewCD(strconv.FormatInt(sum, 10))
}
