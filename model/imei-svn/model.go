package imei_svn_model

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// SVN — Software Version Number.
type SVN string

func (x SVN) IsReserved() bool {
	return x == "99"
}

func (x SVN) Validate() error {
	if len(x) != 2 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return errors.New("invalid SVN")
	}

	return nil
}

func NewSVN(s string) (SVN, error) {
	x := SVN(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}
