package imei_snr

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidSNR = errors.New("invalid SNR")

// SNR — Serial number.
type SNR string

func (x SNR) IsValid() bool {
	return x.Validate() == nil
}

func (x SNR) String() string {
	return string(x)
}

func (x SNR) Validate() error {
	if len(x) != 6 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return ErrInvalidSNR
	}

	return nil
}

func NewSNR(s string) (SNR, error) {
	x := SNR(s)

	if err := x.Validate(); err != nil {
		return "", err
	}

	return x, nil
}
