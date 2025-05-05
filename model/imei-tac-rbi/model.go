package imei_tac_rbi_model

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

//nolint:gochecknoglobals // OK.
var test = map[RBI]struct{}{
	"02": {}, "03": {}, "04": {}, "05": {}, "06": {}, "07": {}, "08": {}, "09": {},
}

//nolint:gochecknoglobals // OK.
var whitelist = map[RBI]struct{}{
	"00": {},
	"01": {},
	"02": {}, "03": {}, "04": {}, "05": {}, "06": {}, "07": {}, "08": {}, "09": {},
	"10": {},
	"30": {},
	"33": {},
	"35": {}, "44": {}, "98": {},
	"45": {},
	"49": {},
	"50": {},
	"51": {},
	"52": {},
	"53": {},
	"54": {},
	"86": {},
	"91": {},
	"99": {},
}

// RBI — Reporting Body Identifier.
//
// https://en.wikipedia.org/wiki/Reporting_Body_Identifier
type RBI string

func (x RBI) IsTestIMEI() bool {
	_, ok := test[x]

	return ok
}

func (x RBI) Validate() error {
	if len(x) != 2 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return errors.New("invalid RBI")
	}

	if _, ok := whitelist[x]; !ok {
		return errors.New("invalid RBI")
	}

	return nil
}

func NewRBI(s string) (RBI, error) {
	x := RBI(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}
