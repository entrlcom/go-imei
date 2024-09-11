package imei_svn

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidSVN = errors.New("invalid SVN")

// SVN — Software Version Number.
type SVN string

func (x SVN) IsReserved() bool {
	return x == "99"
}

func (x SVN) IsValid() bool {
	return x.Validate() == nil
}

func (x SVN) IsZero() bool {
	return len(x) == 0
}

func (x SVN) String() string {
	return string(x)
}

func (x SVN) Validate() error {
	if len(x) != 2 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return ErrInvalidSVN
	}

	return nil
}

func NewSVN(s string) (SVN, error) {
	x := SVN(s)

	if err := x.Validate(); err != nil {
		return "", err
	}

	return x, nil
}
