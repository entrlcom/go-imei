package imei_snr_model

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// SNR — Serial number.
type SNR string

func (x SNR) Validate() error {
	if len(x) != 6 || !strings.ContainsFunc(string(x), unicode.IsDigit) {
		return errors.New("invalid SNR")
	}

	return nil
}

func NewSNR(s string) (SNR, error) {
	x := SNR(s)

	if err := x.Validate(); err != nil {
		return "", fmt.Errorf("x.Validate: %w", err)
	}

	return x, nil
}
