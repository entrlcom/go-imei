package imei_tac

import (
	"errors"
	"strings"
	"unicode"

	imei_tac_rbi "entrlcom.dev/imei/tac/rbi"
)

var ErrInvalidTAC = errors.New("invalid TAC")

// TAC — Type Allocation Code.
//
// https://en.wikipedia.org/wiki/Type_Allocation_Code
type TAC struct {
	id  string
	rbi imei_tac_rbi.RBI
}

func (x TAC) ID() string {
	return x.id
}

func (x TAC) IsValid() bool {
	return x.Validate() == nil
}

func (x TAC) RBI() imei_tac_rbi.RBI {
	return x.rbi
}

func (x TAC) Validate() error {
	if err := x.rbi.Validate(); err != nil {
		return errors.Join(err, ErrInvalidTAC)
	}

	if len(x.id) != 6 || !strings.ContainsFunc(x.id, unicode.IsDigit) {
		return ErrInvalidTAC
	}

	return nil
}

func NewTAC(s string) (TAC, error) {
	if len(s) != 8 { //nolint:gomnd // OK.
		return TAC{}, ErrInvalidTAC
	}

	rbi, err := imei_tac_rbi.NewRBI(s[:2])
	if err != nil {
		return TAC{}, errors.Join(err, ErrInvalidTAC)
	}

	x := TAC{
		id:  s[2:],
		rbi: rbi,
	}

	if err = x.Validate(); err != nil {
		return TAC{}, err
	}

	return x, nil
}
