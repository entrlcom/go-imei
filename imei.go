package imei

import (
	"errors"
	"fmt"
	"strings"
	"unicode"

	imei_cd_model "flida.dev/imei/model/imei-cd"
	imei_snr_model "flida.dev/imei/model/imei-snr"
	imei_svn_model "flida.dev/imei/model/imei-svn"
	"flida.dev/imei/model/imei-tac"
)

// IMEI — International Mobile Equipment Identity.
//
// https://en.wikipedia.org/wiki/International_Mobile_Equipment_Identity
type IMEI struct {
	cd  imei_cd_model.CD
	snr imei_snr_model.SNR
	svn imei_svn_model.SVN
	tac imei_tac_model.TAC
}

func (x IMEI) GetCD() imei_cd_model.CD {
	return x.cd
}

func (x IMEI) GetSNR() imei_snr_model.SNR {
	return x.snr
}

func (x IMEI) GetSVN() imei_svn_model.SVN {
	return x.svn
}

func (x IMEI) GetTAC() imei_tac_model.TAC {
	return x.tac
}

func (x IMEI) IsIMEI() bool {
	return x.cd != ""
}

func (x IMEI) IsIMEISV() bool {
	return x.svn != ""
}

func (x IMEI) String() string {
	if x.IsIMEI() {
		return fmt.Sprintf("%s %s %s %s", string(x.tac.GetRBI()), x.tac.GetID(), string(x.snr), string(x.cd))
	}

	return fmt.Sprintf("%s %s %s %s", string(x.tac.GetRBI()), x.tac.GetID(), string(x.snr), string(x.svn))
}

//nolint:cyclop,gocognit // OK.
func (x IMEI) Validate() error {
	if err := x.tac.Validate(); err != nil {
		return fmt.Errorf("x.imei-tac.Validate: %w", err)
	}

	if err := x.snr.Validate(); err != nil {
		return fmt.Errorf("x.imei-snr.Validate: %w", err)
	}

	if x.IsIMEI() { //nolint:nestif // OK.
		if x.svn != "" {
			return errors.New("invalid IMEI")
		}

		if err := x.cd.Validate(); err != nil {
			return fmt.Errorf("x.imei-cd.Validate: %w", err)
		}

		cd, err := imei_cd_model.ComputeCD(string(x.tac.GetRBI()) + string(x.tac.GetID()) + string(x.snr))
		if err != nil {
			return fmt.Errorf("imei_cd_model.ComputeCD: %w", err)
		}

		if x.cd != cd {
			return errors.New("invalid IMEI")
		}
	}

	if x.IsIMEISV() {
		if x.cd != "" {
			return errors.New("invalid IMEI")
		}

		if err := x.svn.Validate(); err != nil {
			return fmt.Errorf("x.imei-svn.Validate: %w", err)
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
	case 15: //nolint:mnd // OK.
		return newIMEI(s)
	case 16: //nolint:mnd // OK.
		return newIMEISV(s)
	default:
		return IMEI{}, errors.New("invalid IMEI")
	}
}

func newIMEI(s string) (IMEI, error) {
	snr, err := imei_snr_model.NewSNR(s[8:14])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_snr.NewSNR: %w", err)
	}

	tac, err := imei_tac_model.ParseTAC(s[:8])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_tac.ParseTAC: %w", err)
	}

	cd, err := imei_cd_model.NewCD(s[14:15])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_cd.NewCD: %w", err)
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
	snr, err := imei_snr_model.NewSNR(s[8:14])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_snr.NewSNR: %w", err)
	}

	svn, err := imei_svn_model.NewSVN(s[14:16])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_svn.NewSVN: %w", err)
	}

	tac, err := imei_tac_model.ParseTAC(s[:8])
	if err != nil {
		return IMEI{}, fmt.Errorf("imei_tac.ParseTAC: %w", err)
	}

	x := IMEI{
		cd:  "",
		snr: snr,
		svn: svn,
		tac: tac,
	}

	if err = x.Validate(); err != nil {
		return IMEI{}, fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}
