package imei

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	imei_cd "entrlcom.dev/imei/cd"
	imei_snr "entrlcom.dev/imei/snr"
	imei_svn "entrlcom.dev/imei/svn"
	imei_tac "entrlcom.dev/imei/tac"
)

var ErrInvalidIMEI = errors.New("invalid IMEI")

// IMEI — International Mobile Equipment Identity.
//
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity
type IMEI struct {
	cd  imei_cd.CD
	snr imei_snr.SNR
	svn imei_svn.SVN
	tac imei_tac.TAC
}

func (x IMEI) CD() imei_cd.CD {
	return x.cd
}

func (x IMEI) IsIMEI() bool {
	return !x.cd.IsZero()
}

func (x IMEI) IsIMEISV() bool {
	return !x.svn.IsZero()
}

func (x IMEI) SNR() imei_snr.SNR {
	return x.snr
}

func (x IMEI) String() string {
	if x.IsIMEI() {
		return fmt.Sprintf(`%s %s %s %s`, x.tac.RBI().String(), x.tac.ID(), x.snr.String(), x.cd.String())
	}

	return fmt.Sprintf(`%s %s %s %s`, x.tac.RBI().String(), x.tac.ID(), x.snr.String(), x.svn.String())
}

func (x IMEI) SVN() imei_svn.SVN {
	return x.svn
}

func (x IMEI) TAC() imei_tac.TAC {
	return x.tac
}

//nolint:cyclop,gocognit // OK.
func (x IMEI) Validate() error {
	if err := x.tac.Validate(); err != nil {
		return errors.Join(err, ErrInvalidIMEI)
	}

	if err := x.snr.Validate(); err != nil {
		return errors.Join(err, ErrInvalidIMEI)
	}

	if x.IsIMEI() { //nolint:nestif // OK.
		if !x.svn.IsZero() {
			return ErrInvalidIMEI
		}

		if err := x.cd.Validate(); err != nil {
			return errors.Join(err, ErrInvalidIMEI)
		}

		cd, err := imei_cd.ComputeCD(x.tac.RBI().String() + x.tac.ID() + x.snr.String())
		if err != nil {
			return errors.Join(err, ErrInvalidIMEI)
		}

		if !x.cd.IsEqual(cd) {
			return ErrInvalidIMEI
		}
	}

	if x.IsIMEISV() {
		if !x.cd.IsZero() {
			return ErrInvalidIMEI
		}

		if err := x.svn.Validate(); err != nil {
			return errors.Join(err, ErrInvalidIMEI)
		}
	}

	return nil
}

func NewIMEI(s string) (IMEI, error) {
	s = strings.Map(func(r rune) rune {
		if unicode.IsDigit(r) {
			return r
		}

		return -1
	}, s)

	switch len(s) {
	case 15: //nolint:gomnd // OK.
		return newIMEI(s)
	case 16: //nolint:gomnd // OK.
		return newIMEISV(s)
	default:
		return IMEI{}, ErrInvalidIMEI
	}
}

func newIMEI(s string) (IMEI, error) {
	snr, err := imei_snr.NewSNR(s[8:14])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	tac, err := imei_tac.NewTAC(s[:8])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	cd, err := imei_cd.NewCD(s[14:15])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	x := IMEI{
		cd:  cd,
		snr: snr,
		svn: "",
		tac: tac,
	}

	if err = x.Validate(); err != nil {
		return IMEI{}, err
	}

	return x, nil
}

func newIMEISV(s string) (IMEI, error) {
	snr, err := imei_snr.NewSNR(s[8:14])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	svn, err := imei_svn.NewSVN(s[14:16])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	tac, err := imei_tac.NewTAC(s[:8])
	if err != nil {
		return IMEI{}, errors.Join(err, ErrInvalidIMEI)
	}

	x := IMEI{
		cd:  "",
		snr: snr,
		svn: svn,
		tac: tac,
	}

	if err = x.Validate(); err != nil {
		return IMEI{}, err
	}

	return x, nil
}
